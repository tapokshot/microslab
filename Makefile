.PHONY: build
build:
	go build -v ./cmd/mixbot_server

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build