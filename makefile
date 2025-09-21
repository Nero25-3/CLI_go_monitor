APP_NAME=monitor

.PHONY: run lint test clean build tidy start stop status check-json check-html

run:
	go run main.go check https://google.com https://badurl.com -t 5 -l results.log


lint:
	golangci-lint run ./...

test:
	go test -v ./...

clean:
	go clean
	rm -rf $(APP_NAME)
	rm -rf results.log

build:
	go build -o $(APP_NAME) main.go

tidy:
	go mod tidy
	go mod vendor

start:
	go run main.go start $(URLS) -i $(INTERVAL)

stop:
	go run main.go stop

status:
	go run main.go status

clean-cache:
	go clean -modcache

install-deps:
	go mod download
run-yaml:
	go run main.go start -c config.yaml

check-json:
	go run main.go check https://google.com https://example.com --export-json results.json

check-html:
	go run main.go check https://google.com https://example.com --export-html results.html

check-all:
	go run main.go check https://google.com https://example.com --export-json results.json --export-html results.html