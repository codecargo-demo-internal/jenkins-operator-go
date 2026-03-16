.PHONY: build test lint clean

build:
	go build -o bin/operator ./cmd/...

test:
	go test -v -coverprofile=coverage.out ./...

lint:
	golangci-lint run ./...

clean:
	rm -rf bin/ coverage.out
