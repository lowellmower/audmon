package cli

import (
    "github.com/lowellmower/audmon/pkg/config"
    "github.com/lowellmower/audmon/pkg/daemon"
    "github.com/spf13/cobra"
)

var (
    foreground bool
)

func init() {
    serverCmd.Flags().BoolVar(&foreground, "foreground", false, "--foreground=true")
    //serverCmd.Flags().StringVar(&srvAddr, "address", "0.0.0.0", "--address=<IPv4 address>")
    //serverCmd.Flags().StringVar(&srvPort, "port", "9009", "--address=<IPv4 address>")
    rootCmd.AddCommand(daemonCmd)
}

// daemonCmd starts the audmon daemon
var daemonCmd = &cobra.Command{
    Use:   "daemon",
    Short: "Starts the audmon daemon.",
    Long: `Some words about the audmon daemon...`,
    RunE: func(cmd *cobra.Command, args []string) error {
        da := config.DaemonArgs{Foreground: foreground}
        config.AppConf.Daemon = da
        return daemon.Run()
    },
}
