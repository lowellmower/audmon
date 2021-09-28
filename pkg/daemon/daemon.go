package daemon

import (
	"flag"
	"github.com/lowellmower/audmon/pkg/config"
	"github.com/pkg/errors"
	"io"
	"os"

	"github.com/lowellmower/audmon/pkg/log"

	"github.com/elastic/go-libaudit/v2"
	"github.com/elastic/go-libaudit/v2/auparse"
)

var (
	fs          = flag.NewFlagSet("audit", flag.ExitOnError)
	diag        = fs.String("diag", "", "dump raw information from kernel to file")
	rate        = fs.Uint("rate", 0, "rate limit in kernel (default 0, no rate limit)")
	backlog     = fs.Uint("backlog", 8192, "backlog limit")
	immutable   = fs.Bool("immutable", false, "make kernel audit settings immutable (requires reboot to undo)")
	receiveOnly = fs.Bool("ro", false, "receive only using multicast, requires kernel 3.16+")
)

func Run() error {
	if config.AppConf.Daemon.Foreground {
		return read()
	}
	fs.Parse(os.Args[1:])

	return read()
}

func read() error {
	var diagWriter io.Writer
	client, err := libaudit.NewAuditClient(diagWriter)
	if err != nil {
		return err
	}

	defer client.Close()

	status, err := client.GetStatus()
	if err != nil {
		return errors.Wrap(err, "failed to get audit status")
	}

	log.Infof("received audit status=%+v", status)

	if status.Enabled == 0 {
		log.Info("enabling auditing in the kernel")
		if err = client.SetEnabled(true, libaudit.WaitForReply); err != nil {
			return err
		}
	}

	if status.RateLimit != uint32(*rate) {
		log.Infof("setting rate limit in kernel to %v", *rate)
		if err = client.SetRateLimit(uint32(*rate), libaudit.NoWait); err != nil {
			return errors.Wrap(err, "failed to set rate limit to unlimited")
		}
	}

	if status.BacklogLimit != uint32(*backlog) {
		log.Infof("setting backlog limit in kernel to %v", *backlog)
		if err = client.SetBacklogLimit(uint32(*backlog), libaudit.NoWait); err != nil {
			return errors.Wrap(err, "failed to set backlog limit")
		}
	}

	if status.Enabled != 2 {
		log.Infof("setting kernel settings as immutable")
		if err = client.SetImmutable(libaudit.NoWait); err != nil {
			return errors.Wrap(err, "failed to set kernel as immutable")
		}
	}

	log.Infof("sending message to kernel registering our PID (%v) as the audit daemon", os.Getpid())
	if err = client.SetPID(libaudit.NoWait); err != nil {
		return errors.Wrap(err, "failed to set audit PID")
	}

	return receive(client)
}

func receive(c *libaudit.AuditClient) error {
	for {
		event, err := c.Receive(false)
		if err != nil {
			return err
		}

		if event.Type < auparse.AUDIT_USER_AUTH || event.Type > auparse.AUDIT_LAST_USER_MSG2 {
			continue
		}

		log.Infof("type=%v msg=%v", event.Type, string(event.Data))
	}
}
