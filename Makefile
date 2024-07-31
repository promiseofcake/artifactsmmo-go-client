TOOLS_PATH := $(shell pwd)/bin

install-tools:
	export GOBIN=$(TOOLS_PATH); cat tools.go | grep _ | awk -F'("|//)' '{print $$NF " " $$2}' | xargs -tL 1 go install

generate-client:
	go generate ./...
