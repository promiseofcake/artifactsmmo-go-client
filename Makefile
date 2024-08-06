TOOLS_PATH := $(shell pwd)/bin

install-tools:
	export GOBIN=$(TOOLS_PATH); cat tools.go | grep _ | awk -F'("|//)' '{print $$NF " " $$2}' | xargs -tL 1 go install

fetch-spec:
	curl --silent "https://api.artifactsmmo.com/openapi.json" | jq . > spec/openapi.json

all: install-tools fetch-spec process-spec generate-client test

process-spec:
	openapi-down-convert --input spec/openapi.json --output spec/openapi-3.0.json
	./scripts/format-spec.sh spec/openapi-3.0.json

generate-client:
	go generate ./...

test:
	go test -v ./...
