package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the running monitor process",
	RunE: func(cmd *cobra.Command, args []string) error {
		pidFile := "monitor.pid"
		data, err := os.ReadFile(pidFile)
		if err != nil {
			return errors.New("monitor is not running")
		}
		pid, err := strconv.Atoi(string(data))
		if err != nil {
			return errors.New("invalid PID file")
		}
		process, err := os.FindProcess(pid)
		if err != nil {
			return errors.New("process not found")
		}
		if err := process.Signal(syscall.SIGTERM); err != nil {
			return fmt.Errorf("failed to stop process: %w", err)
		}
		if err := os.Remove(pidFile); err != nil {
			return fmt.Errorf("failed to remove PID file: %w", err)
		}
		fmt.Println("Monitor stopped.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
