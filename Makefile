MAIN_PKG := github.com/zendesk/melbourne_code_club_go/cmd/melbourne_code_club_go
COVERAGE_FILE_NAME := coverage.out
TEST_TARGET ?= ./...
TEST_FLAGS := -race -coverpkg=./...  --coverprofile=$(COVERAGE_FILE_NAME)

default: build

run:
	go run cmd/melbourne_code_club_go/main.go

build:
	go build -race "$(MAIN_PKG)"

ensure_deps:
	go mod tidy
	go mod vendor

lint:
	$(GOPATH)/bin/golangci-lint run ./...

lint_fix:
	$(GOPATH)/bin/golangci-lint run --fix ./...

install_devtools:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin $(GOLANGCI_VERSION)
	go install gotest.tools/gotestsum
	go install github.com/go-delve/delve/cmd/dlv
	go install github.com/githubnemo/CompileDaemon

test:
	CGO_ENABLED=1 gotestsum -- $(TEST_FLAGS) $(TEST_TARGET) -timeout 5s