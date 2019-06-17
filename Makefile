.DEFAULT_GOAL := default
.EXPORT_ALL_VARIABLES: 
GO111MODULE=on

default: build

build: build-example build-gen

build-example: cmd/example/main.go
	@go build ./cmd/example

build-gen: cmd/gen/main.go
	@go build ./cmd/gen

generate-swagger-json:
	@pip install pyyaml pyyaml-include
	@python scripts/parseincludes.py > schema/compiled.yaml
	@npm install -g api-spec-converter
	@api-spec-converter schema/compiled.yaml --from=openapi_3 --to=swagger_2 -d > schema/swagger.json

generate-models:
	@go get -u github.com/go-swagger/go-swagger/cmd/swagger
	@swagger generate model -f schema/swagger.json

generate-client:
	@go get -u github.com/go-swagger/go-swagger/cmd/swagger
	@swagger generate client -f schema/swagger.json

generate-methods: build-gen
	@$(PWD)/gen > rpc_subscriptions.go
	goimports -w rpc_subscriptions.go
	@gofmt -w rpc_subscriptions.go

.PHONY: generate-models generate-client generate-methods generate-swagger-json