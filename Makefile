.DEFAULT_GOAL := default

default: generate-models generate-methods

generate-models:
	@go get github.com/ChimeraCoder/gojson/gojson
	@scripts/generate-models.sh

generate-methods:
	@scripts/generate-methods.sh

.PHONY: generate-models generate-methods