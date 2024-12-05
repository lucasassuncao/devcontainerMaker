.PHONY: run install test docs lint build release

default: run

run:
	go run .

install:
	go build

test:
	go test -v ./...

docs:
	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
	gomarkdoc -e -o '{{.Dir}}/README.md' ./...

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint -v run ./...

build:
	go install github.com/goreleaser/goreleaser@latest
	goreleaser build --skip-validate --single-target --snapshot --clean

release:
	go install github.com/goreleaser/goreleaser@latest
	goreleaser release --timeout 360s