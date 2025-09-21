package cmd

import (
	"fmt"
	"sync"
	"time"

	"CLI_go_monitor/internal"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var timeout int
var logFilePath string

type result struct {
	url string
	err error
}

var checkCmd = &cobra.Command{
	Use:   "check [url1] [url2] ... ",
	Short: "Check the status of one or more URLs",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logger := internal.NewLogger(logFilePath)

		var wg sync.WaitGroup
		results := make(chan result, len(args))

		for _, url := range args {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				err := internal.MonitorURL(url, time.Duration(timeout)*time.Second)
				results <- result{url, err}
			}(url)
		}
		wg.Wait()
		close(results)

		successCount, failCount := 0, 0
		for r := range results {
			if r.err == nil {
				msg := fmt.Sprintf("%-40s [OK]", r.url)
				color.New(color.FgGreen).Println(msg)
				logger.Info(msg)
				successCount++
			} else {
				msg := fmt.Sprintf("%-40s [FAIL] (%v)", r.url, r.err)
				color.New(color.FgRed).Println(msg)
				logger.Info(msg)
				failCount++
			}
		}

		summary := fmt.Sprintf("\nSummary: %d successful, %d failed\n", successCount, failCount)
		color.New(color.FgCyan, color.Bold).Println(summary)
		logger.Info(summary)
	},
}

func init() {
	checkCmd.Flags().IntVarP(&timeout, "timeout", "t", 5, "Timeout in seconds")
	checkCmd.Flags().StringVarP(&logFilePath, "logfile", "l", "monitor.log", "Path to log file")
	rootCmd.AddCommand(checkCmd)
}
