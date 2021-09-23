package cli

import (
	"github.com/lowellmower/audmon/pkg/client"

	"github.com/spf13/cobra"
)

var (
	clientAddr string
	clientPort string
)

func init() {
	clientCmd.Flags().StringVar(&clientAddr, "address", "0.0.0.0", "--address=<IPv4 address>")
	clientCmd.Flags().StringVar(&clientPort, "port", "9009", "--address=<IPv4 address>")
	rootCmd.AddCommand(clientCmd)
}

// clientCmd starts the audmon client capable of receiving messages from the
// audmon server should all firewalls and ingress rules on a host allow.
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Starts the audmon client.",
	Long: `Some words about the audmon client...`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return client.Run(clientAddr+":"+clientPort)
	},
}
