# CLI_go_monitor
Basic CLI URL monitoring tool built with Go and the Cobra library. Performs HTTP checks on URLs with configurable timeout, reporting their status and availability. Designed as a starter project to expand with concurrency, statistics, and logging features.

## Architecture

The application follows a modular command-based structure using the Cobra package:

"""

CLI_go_monitor/
├── main.go # Application entry point & command setup
├── cmd/ # Package containing CLI commands
│ ├── root.go # Root command (base CLI)
│ └── check.go # 'check' command to monitor individual URLs
├── internal/ # Internal packages (future expansions)
│ ├── monitor.go # URL monitoring logic and helpers
│ └── utils.go # Utility functions (e.g., formatting output)
└── go.mod # Go module file

"""

