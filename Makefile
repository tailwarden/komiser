# Build all
VERSION := $(shell git rev-parse --abbrev-ref HEAD)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
COMMIT := $(shell git rev-parse --short HEAD)

EXECUTABLE=komiser

GO_LD_FLAGS += -X github.com/tailwarden/komiser/internal.Version=$(VERSION)
GO_LD_FLAGS += -X github.com/tailwarden/komiser/internal.Buildtime=$(BUILDTIME)
GO_LD_FLAGS += -X github.com/tailwarden/komiser/internal.Commit=$(COMMIT)
GO_FLAGS = -ldflags "$(GO_LD_FLAGS)"

build: ## Build for the current platform
	go build -o bin/$(EXECUTABLE) $(GO_FLAGS) .
	@echo built: bin/$(EXECUTABLE)
	@echo version: $(VERSION)
	@echo commit: $(COMMIT)

package: ## Build for all platforms
	env GOOS=windows GOARCH=amd64 go build -o bin/$(EXECUTABLE)_windows_amd64.exe $(GO_FLAGS) .
	env GOOS=linux GOARCH=amd64 go build -o bin/$(EXECUTABLE)_linux_amd64 $(GO_FLAGS) .
	env GOOS=darwin GOARCH=amd64 go build -o bin/$(EXECUTABLE)_darwin_amd64 $(GO_FLAGS) .
	env GOOS=darwin GOARCH=arm64 go build -o bin/$(EXECUTABLE)_darwin_arm64 $(GO_FLAGS) .
	@echo built:  bin/$(EXECUTABLE)_windows_amd64.exe, bin/$(EXECUTABLE)_linux_amd64, bin/$(EXECUTABLE)_darwin_amd64, bin/$(EXECUTABLE)_darwin_arm64
	@echo version: $(VERSION)
	@echo commit: $(COMMIT)

test: ## Run tests
	go test -v ./...

version:
	@echo version: $(VERSION)

clean: ## Remove previous builds and clear test cache
	rm -f bin/$(EXECUTABLE)*
	go clean -testcache