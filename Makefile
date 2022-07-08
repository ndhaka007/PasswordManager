MIGRATION_OUT := "build/migration/migration"
MIGRATION_PKG_BUILD := "cmd/migration/main.go"

BUF_VERSION := v0.42.1
PROTOC_GEN_GO_VERSION := v1.3.2
PROTOC_GEN_GRPC_GATEWAY_VERSION := v1.14.3

RPC_ROOT := rpc/

UNAME_OS=$(shell go env GOOS)
UNAME_ARCH=$(shell go env GOARCH)

.PHONY: init ## First time setup
init:
	cp config/env.default.toml config/env.dev.toml

.PHONY: deps ## Download the dependencies
deps:
	@echo "\n + Downloading dependencies \n"
	@go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)
	@go install github.com/golang/protobuf/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@$(PROTOC_GEN_GRPC_GATEWAY_VERSION)
	@go mod vendor

.PHONY: clean ## Remove previous builds, protobuf files, and proto compiled code
clean:
	@echo " + Removing cloned and generated files\n"
	@rm -rf $(MIGRATION_OUT) $(RPC_ROOT)

.PHONY: proto-generate ## Compile protobuf to pb
proto-generate:
	@echo "\n + Generating pb language bindings\n"
	@buf ls-files
	@buf generate

.PHONY: proto-refresh ## Download and re-compile protobuf
proto-refresh: clean proto-generate

.PHONY: go-build-migration ## Build the binary file for the migration
go-build-migration:
	@CGO_ENABLED=0 GOOS=$(UNAME_OS) GOARCH=$(UNAME_ARCH) go build -i -v -o $(MIGRATION_OUT) $(MIGRATION_PKG_BUILD)

.PHONY: go-run-migration ## Run the migration
go-run-migration:
	@go run $(MIGRATION_PKG_BUILD) up

.PHONY: help ## Display this help screen
help:
	@echo "Usage:"
	@grep -E '^\.PHONY: [a-zA-Z_-]+.*?## .*$$' $(MAKEFILE_LIST) | sort | sed 's/\.PHONY\: //' | awk 'BEGIN {FS = " ## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'