.PHONY: install test

default: run

run:
	go run .

install:
	go build

test:
	go test

docs:
	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
	gomarkdoc -e -o '{{.Dir}}/README.md' ./...

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint -v run ./...
