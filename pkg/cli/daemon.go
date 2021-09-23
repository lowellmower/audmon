package cli

import (
	"github.com/lowellmower/audmon/pkg/daemon"
    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(daemonCmd)
}

// daemonCmd starts the audmon daemon
var daemonCmd = &cobra.Command{
    Use:   "daemon",
    Short: "Starts the audmon daemon.",
    Long: `Some words about the audmon daemon...`,
    RunE: func(cmd *cobra.Command, args []string) error {
        daemon.Printer()
        return nil
    },
}
