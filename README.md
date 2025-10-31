# CLI_go_monitor
Basic CLI URL monitoring tool built with Go and the Cobra library. Performs HTTP checks on URLs with configurable timeout, reporting their status and availability. Designed as a starter project to expand with concurrency, statistics, and logging features.

## CLI_go_monitor

Advanced Endpoint Monitoring – Easy Configuration, Extensible Architecture

---

## Why this project?

In my professional experience supporting critical systems, most monitoring tools lack flexibility, security, and ease of integration. Teams often rely on monolithic scripts that are tough to maintain, poorly tested, and restricted in alerting capabilities.

**CLI_go_monitor** is my response, combining:

- Strong input validation
- Modular, plugin-oriented architecture
- Multi-channel alerts (Slack, Telegram, Email, Webhooks)
- Security by design: protected logs, process isolation

---

## Key Features

- Configuration via YAML, JSON, TOML, or ENV
- Event system and plugins for easy API/service integration
- Advanced alerting: throttling, grouping, custom alert channels  
- Built-in performance metrics (Prometheus endpoint)
- CLI with autocompletion and clear help (Cobra)
- Unit, integration, and fuzz tests included
- Dockerfile for fast deployment
- 100% open source and easily extensible

---

## System Architecture

The following diagram illustrates the high-level architecture and data flows for CLI_go_monitor:


![System Architecture Diagram](/docs/CLI_go_monitor.png)

## Project structure

```
CLI_go_monitor/
│
├── cmd/                     # CLI entrypoint (main.go and commands)
│   └── main.go
│
├── internal/                # Core logic (not externally importable)
│   ├── core/                # Scheduler, event hub, plugin orchestrator
│   ├── config/              # Parsing, validation, config management
│   ├── plugins/             # Interfaces and plugin registry
│   │   ├── check.go         # Interface for monitor plugins
│   │   └── notifier.go      # Interface for alert plugins
│   ├── monitor/             # Implementation of monitor plugins (HTTP, ping, etc.)
│   └── alert/               # Implementation of alert plugins (Slack, email, etc.)
│
├── pkg/                     # Reusable libraries/utilities (if needed)
│
├── test/                    # Centralized integration/acceptance tests
│
├── docs/                    # Documentation, diagrams, examples
│   └── architecture.md
│
├── scripts/                 # Utility scripts for development, CI/CD, etc.
│
├── Dockerfile               # Container setup for easy deployment
├── README.md
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
└── CHANGELOG.md
```


**Description:** 
- **CLI Entry (`main.go`):** Parses user arguments, loads configuration, serves as the application's entry point.
- **Configuration:** Supports YAML, JSON, TOML, ENV files for flexible setup.
- **Core Logic:** Schedules checks, manages event flow, and orchestrates plugins.
- **Monitor Plugins:** Extensible check modules for HTTP, ping, database, etc.
- **Alert Plugins:** Extensible alerting for Slack, Email, Telegram, webhooks.
- **Results Logging/Storage:** Collects and stores results, exposes metrics for Prometheus, handles persistent logging.

---

## Example Usage

cli_go_monitor --config=config.yaml check --verbose


*Check the `/docs` directory for demo GIFs and screenshots.*

---

## How to Contribute

Open to PRs, suggestions, and feedback:
1. Fork the repository.
2. Install dependencies and run tests locally.
3. Review the ROADMAP and Issues for top priorities.

---

## Differentiation

Compared to XMonitor, YWatchdog, and others, CLI_go_monitor emphasizes:

- **Security:** Robust validation and protected logging
- **Flexibility:** Plugin system & multi-config support
- **Professionalism:** Testing and CI/CD from day one

Upcoming features on the roadmap: ML-based downtime prediction and real-time dashboard.

---

## License  
MIT – Free for personal and professional use.

## Contact  
[brunocoya25@gmail.com] – [linkedin.com/in/bruno-coya-030470184]
