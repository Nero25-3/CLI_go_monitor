APP_NAME=monitor

.PHONY: run lint test clean

run:
	go run main.go

lint:
	golangci-lint run ./...

test:
	go test -v ./...

clean:
	go clean
