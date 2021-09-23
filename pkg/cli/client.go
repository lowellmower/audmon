package cli

import (
	"github.com/lowellmower/audmon/pkg/client"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clientCmd)
}

// clientCmd starts the audmon client capable of receiving messages from the
// audmon server should all firewalls and ingress rules on a host allow.
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Starts the audmon client.",
	Long: `Some words about the audmon client...`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client.Printer()
		return nil
	},
}
