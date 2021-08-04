MAIN_PKG := github.com/zendesk/melbourne_code_club_go/cmd/melbourne_code_club_go

default: build

build:
	go build "$(MAIN_PKG)"

ensure_deps:
	go mod tidy
	go mod vendor

lint:
	$(GOPATH)/bin/golangci-lint run ./...

lint_fix:
	$(GOPATH)/bin/golangci-lint run --fix ./...