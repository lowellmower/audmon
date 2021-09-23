package cli

import (
    "github.com/lowellmower/audmon/pkg/server"
    "github.com/spf13/cobra"
)

var (
    srvAddr string
    srvPort string
)

func init() {
    serverCmd.Flags().StringVar(&srvAddr, "address", "0.0.0.0", "--address=<IPv4 address>")
    serverCmd.Flags().StringVar(&srvPort, "port", "9009", "--address=<IPv4 address>")
    rootCmd.AddCommand(serverCmd)
}

// serverCmd starts the audmon server for communication with a client
var serverCmd = &cobra.Command{
    Use:   "server",
    Short: "Starts the audmon server.",
    Long: `Some words about the audmon server...`,
    RunE: func(cmd *cobra.Command, args []string) error {
        return server.Run(srvAddr+":"+srvPort)
    },
}

