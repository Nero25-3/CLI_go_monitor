package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"CLI_go_monitor/internal"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var configPath string
var interval int
var pidFile = "monitor.pid"

var startCmd = &cobra.Command{
	Use:   "start [url1] [url2] ...",
	Short: "Start periodic monitoring of URLs",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat(pidFile); err == nil {
			return errors.New("monitor is already running")
		}

		var urls []string
		var timeout int
		var logfile string

		if configPath != "" {
			cfg, err := internal.ReadConfig(configPath)
			if err != nil {
				return fmt.Errorf("failed to read config file: %w", err)
			}
			urls = cfg.URLs
			interval = cfg.Interval
			timeout = cfg.Timeout
			logfile = cfg.Logfile
		} else {
			urls = args
			timeout = 5
			logfile = "monitor.log"
		}

		if len(urls) == 0 {
			return errors.New("no URLs provided")
		}

		logger := internal.NewLogger(logfile)

		go func() {
			ticker := time.NewTicker(time.Duration(interval) * time.Second)
			defer ticker.Stop()

			var wg sync.WaitGroup

			for range ticker.C {
				wg.Add(len(urls))
				for _, url := range urls {
					go func(u string) {
						defer wg.Done()
						err := internal.MonitorURL(u, time.Duration(timeout)*time.Second)
						if err == nil {
							msg := fmt.Sprintf("[OK] %s", u)
							color.New(color.FgGreen).Println(msg)
							logger.Info(msg)
						} else {
							msg := fmt.Sprintf("[FAIL] %s - %v", u, err)
							color.New(color.FgRed).Println(msg)
							logger.Info(msg)
						}
					}(url)
				}
				wg.Wait()
			}
		}()

		pid := os.Getpid()
		if err := os.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0644); err != nil {
			return fmt.Errorf("error writing PID file: %w", err)
		}
		fmt.Println("Monitor started with PID", pid)

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		fmt.Println("Stopping monitor...")
		os.Remove(pidFile)
		os.Exit(0)

		return nil
	},
}

func init() {
	startCmd.Flags().StringVarP(&configPath, "config", "c", "", "Path to YAML config file")
	startCmd.Flags().IntVarP(&interval, "interval", "i", 60, "Monitoring interval in seconds (ignored if config is set)")
	rootCmd.AddCommand(startCmd)
}
