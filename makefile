APP_NAME=monitor
DEBUG_PORT=2345
DEBUG_HEADLESS=false

.PHONY: run lint test clean build tidy start stop status check-json check-html debug

run:
	go run cmd/main.go check https://www.google.com https://www.badurl.com -t 5 -l results.log
	
lint:
	golangci-lint run ./...

test:
	go test -v ./...

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
