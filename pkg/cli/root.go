package cli

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

func init() {
    cobra.OnInitialize()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "audmon",
    Short: "An auditd message monitor.",
    Long: `Some words about audmon here...'`,
    // Uncomment the following line if your bare application
    // has an action associated with it:
    //      Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Run() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
