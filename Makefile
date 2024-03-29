all: format test lint build

install:
	go build -o $(GOPATH)/bin/gocopy

format:
	gofmt -w .

test:
	go test -cover ./...

lint:
	bin/golangci-lint run --enable-all

build:
	go build -o bin/gocopy