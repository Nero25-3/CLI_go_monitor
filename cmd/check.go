package cmd

import (
	"CLI_go_monitor/internal"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var timeout int

var checkCmd = &cobra.Command{
	Use:   "check [url]",
	Short: "Check the status of a URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		err := internal.MonitorURL(url, time.Duration(timeout)*time.Second)
		if err != nil {
			fmt.Printf("%s is not reachable: %v\n", url, err)

		} else {
			fmt.Printf("%s is alive\n", url)
		}
	},
}

// init initializes the check command and adds it to the root command.
func init() {
	checkCmd.Flags().IntVarP(&timeout, "timeout", "t", 5, "Timeout in seconds")
	rootCmd.AddCommand(checkCmd)
}
