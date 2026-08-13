.PHONY: all build build-cross test vet fmt lint cover install tidy clean help

BINARY   := ry
MAIN_PKG := ./cmd/ry
GO       := go
GOFLAGS  :=
OUTPUT   := bin/$(BINARY)

VERSION  := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT   := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
LDFLAGS  := -X main.version=$(VERSION) -X main.commit=$(COMMIT)
BUILD_FLAGS := -ldflags "$(LDFLAGS)"

all: build

build: ## Build the CLI binary into bin/
	$(GO) build $(GOFLAGS) $(BUILD_FLAGS) -o $(OUTPUT) $(MAIN_PKG)

CROSS_GOOS := linux windows darwin
CROSS_GOARCH := amd64 arm64

build-cross: ## Cross-compile for linux/windows/darwin x amd64/arm64 into bin/
	@mkdir -p bin
	@for os in $(CROSS_GOOS); do \
		for arch in $(CROSS_GOARCH); do \
			ext=""; [ "$$os" = "windows" ] && ext=".exe"; \
			echo "Building $$os/$$arch..."; \
			GOOS=$$os GOARCH=$$arch $(GO) build $(GOFLAGS) $(BUILD_FLAGS) -o bin/$(BINARY)-$$os-$$arch$$ext $(MAIN_PKG) || exit 1; \
		done; \
	done

install: ## Install the CLI binary to GOPATH/bin
	$(GO) install $(GOFLAGS) $(BUILD_FLAGS) $(MAIN_PKG)

test: ## Run all tests
	$(GO) test $(GOFLAGS) -race ./...

vet: ## Run go vet
	$(GO) vet ./...

fmt: ## Format all Go source files
	gofmt -s -w .
	$(GO) mod tidy

cover: ## Run tests with coverage and generate HTML report
	$(GO) test $(GOFLAGS) -race -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html

tidy: ## Run go mod tidy
	$(GO) mod tidy

clean: ## Remove build artifacts and coverage files
	rm -rf bin/ coverage.out coverage.html

help: ## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'
