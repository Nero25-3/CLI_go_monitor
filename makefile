APP_NAME=monitor
DEBUG_PORT=2345
DEBUG_HEADLESS=false

.PHONY: run lint clean build tidy start stop status check-json check-html debug

run:
	go run cmd/main.go check https://www.google.com https://www.badurl.com -t 5 -l results.log
	
lint:
	golangci-lint run ./...

clean:
	go clean
	rm -rf $(APP_NAME)
	rm -rf results.log

build:
	go build -o $(APP_NAME) cmd/main.go

tidy:
	go mod tidy
	go mod vendor

start:
	go run cmd/main.go start $(URLS) -i $(INTERVAL)

stop:
	go run cmd/main.go stop

status:
	go run cmd/main.go status

clean-cache:
	go clean -modcache

install-deps:
	go mod download
run-yaml:
	go run cmd/main.go start -c ./config.yaml

check-json:
	go run cmd/main.go check https://google.com https://example.com --export-json results.json

check-html:
	go run cmd/main.go check https://google.com https://example.com --export-html results.html

check-all:
	go run cmd/main.go check https://google.com https://example.com --export-json results.json --export-html results.html

debug:	
	dlv debug --headless=$(DEBUG_HEADLESS) --listen=:$${DEBUG_PORT} --api-version=2 --accept-multiclient ./cmd/main.go -- $(ARGS)

.PHONY: test test-verbose coverage coverage-html coverage-report

# Run tests
test:
	@echo "Running tests..."
	go test ./... -race -timeout 30s

# Run tests with verbose output
test-verbose:
	@echo "Running tests (verbose)..."
	go test ./... -v -race -timeout 30s

# Generate coverage report
coverage:
	@echo "Generating coverage report..."
	go test ./... -cover -coverprofile=coverage.out

# View coverage in HTML
coverage-html: coverage
	@echo "Opening coverage report in browser..."
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Display coverage by function
coverage-report: coverage
	@echo "Coverage by function:"
	@go tool cover -func=coverage.out

# Run all quality checks
quality: test coverage-report lint
	@echo "✅ Quality checks complete"
