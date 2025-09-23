package cmd

import (
	"CLI_go_monitor/internal"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var timeout int
var logFilePath string

var exportJSON string
var exportHTML string

// MonitorURLFunc is a variable to allow mocking in tests
var MonitorURLFunc = internal.MonitorURL

// CheckResult holds the result of checking a URL
type CheckResult struct {
	URL    string `json:"url"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func writeJSON(filename string, results []CheckResult) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(results)
}

func writeHTML(filename string, results []CheckResult) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("<html><body><h1>Check Results</h1><table border='1'><tr><th>URL</th><th>Status</th><th>Error</th></tr>")
	if err != nil {
		return fmt.Errorf("failed to write to HTML file: %w", err)
	}
	for _, r := range results {
		_, err = file.WriteString("<tr>")
		if err != nil {
			return fmt.Errorf("failed to write to HTML file: %w", err)
		}
		_, err = file.WriteString("<td>" + r.URL + "</td>")
		if err != nil {
			return fmt.Errorf("failed to write to HTML file: %w", err)
		}
		_, err = file.WriteString("<td>" + r.Status + "</td>")
		if err != nil {
			return fmt.Errorf("failed to write to HTML file: %w", err)
		}
		if r.Error == "" {
			file.WriteString("<td>-</td>")
		} else {
			file.WriteString("<td>" + r.Error + "</td>")
		}
		file.WriteString("</tr>")
	}
	file.WriteString("</table></body></html>")
	return nil
}

// RunCheck performs the URL checks and handles logging and exporting results
func RunCheck(args []string, timeout int, logfile string, exportJSON string, exportHTML string) error {
	if len(args) == 0 {
		return errors.New("no URLs provided")
	}

	logger := internal.NewLogger(logfile, parseLogLevel(logLevelStr), logEnable)

	var wg sync.WaitGroup
	resultsChan := make(chan CheckResult, len(args))

	for _, url := range args {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			err := MonitorURLFunc(url, time.Duration(timeout)*time.Second)
			var status string
			var errMsg string
			if err == nil {
				status = "OK"
				errMsg = ""
			} else {
				status = "FAIL"
				errMsg = err.Error()
			}
			resultsChan <- CheckResult{URL: url, Status: status, Error: errMsg}

			msg := fmt.Sprintf("%-40s [%s]", url, status)
			if status == "OK" {
				color.New(color.FgGreen).Println(msg)
				logger.Info(msg)
			} else {
				color.New(color.FgRed).Println(msg)
				logger.Error(msg + " " + errMsg)
			}
		}(url)
	}

	wg.Wait()
	close(resultsChan)

	results := []CheckResult{}
	successCount, failCount := 0, 0
	for r := range resultsChan {
		results = append(results, r)
		if r.Status == "OK" {
			successCount++
		} else {
			failCount++
		}
	}

	summary := fmt.Sprintf("\nSummary: %d successful, %d failed\n", successCount, failCount)
	color.New(color.FgCyan, color.Bold).Println(summary)
	logger.Info(summary)

	if exportJSON != "" {
		if err := writeJSON(exportJSON, results); err != nil {
			fmt.Printf("Error writing JSON export: %v\n", err)
		}
	}

	if exportHTML != "" {
		if err := writeHTML(exportHTML, results); err != nil {
			fmt.Printf("Error writing HTML export: %v\n", err)
		}
	}

	return nil
}

var checkCmd = &cobra.Command{
	Use:   "check [url1] [url2] ... ",
	Short: "Check the status of one or more URLs",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunCheck(args, timeout, logFilePath, exportJSON, exportHTML)
	},
}

func init() {
	checkCmd.Flags().IntVarP(&timeout, "timeout", "t", 5, "Timeout in seconds")
	checkCmd.Flags().StringVarP(&logFilePath, "logfile", "l", "monitor.log", "Path to log file")
	checkCmd.Flags().BoolVarP(&logEnable, "log", "g", true, "Enable logging to file")
	checkCmd.Flags().StringVarP(&logLevelStr, "loglevel", "v", "info", "Log level: info, warn, error")
	checkCmd.Flags().StringVarP(&exportJSON, "export-json", "j", "", "File to export results in JSON")
	checkCmd.Flags().StringVarP(&exportHTML, "export-html", "x", "", "File to export results in HTML")

	rootCmd.AddCommand(checkCmd)
}
