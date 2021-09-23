package cli

import (
    "github.com/lowellmower/audmon/pkg/server"
    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(serverCmd)
}

// serverCmd starts the audmon server for communication with a client
var serverCmd = &cobra.Command{
    Use:   "server",
    Short: "Starts the audmon server.",
    Long: `Some words about the audmon server...`,
    RunE: func(cmd *cobra.Command, args []string) error {
        server.Printer()
        return nil
    },
}
