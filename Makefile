.DEFAULT_GOAL := default

default: build

build: cmd/example/main.go
	@go build ./cmd/example

generate-models:
	@go get github.com/ChimeraCoder/gojson/gojson
	@scripts/generate-models.sh

generate-methods:
	@go get golang.org/x/tools/cmd/goimports
	@scripts/generate-methods.sh
	@goimports -w rpc_methods.go
	@gofmt -w rpc_methods.go
	@goimports -w rpc_subscriptions.go
	@gofmt -w rpc_subscriptions.go

.PHONY: generate-models generate-methods