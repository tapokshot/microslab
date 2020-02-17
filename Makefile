.PHONY: build
build:
	go build -v ./cmd/microslab

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: proto
proto:
	protoc --proto_path=./proto ./proto/*.proto --go_out=plugins=grpc:./internal/app/grpc
.DEFAULT_GOAL := build