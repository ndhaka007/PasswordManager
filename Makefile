MIGRATION_OUT := "build/migration/migration"
MIGRATION_PKG_BUILD := "cmd/migration/main.go"

UNAME_OS=$(shell go env GOOS)
UNAME_ARCH=$(shell go env GOARCH)

init:  ## First time setup
	cp config/env.default.toml config/env.dev.toml

deps: ## Download the dependencies
	@go mod vendor

go-build-migration: ## Build the binary file for the migration
	@CGO_ENABLED=0 GOOS=$(UNAME_OS) GOARCH=$(UNAME_ARCH) go build -i -v -o $(MIGRATION_OUT) $(MIGRATION_PKG_BUILD)

go-run-migration: ## Run the migration
	@go run $(MIGRATION_PKG_BUILD) up

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
