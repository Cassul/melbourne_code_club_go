.PHONY: ensure_deps test lint install_devtools get
export GOPROXY := https://$(ARTIFACTORY_USERNAME):$(ARTIFACTORY_API_KEY)@zdrepo.jfrog.io/zdrepo/api/go/zen-go
export GOSUMDB := off

default: test lint

ensure_deps:
	go mod vendor
	go mod tidy

test:
	go test -mod=vendor ./...

lint:
	golangci-lint run ./...

lint_fix:
	golangci-lint run --fix  ./...

install_devtools:
	# as per the instructions on their repo
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.28.3

# -mod=readonly overrides -mod=vendor from GOFLAGS
# The "go get" command remains permitted to update go.mod even with -mod=readonly
get:
ifeq ($(DEP),)
	$(error Please provide the "DEP" variable when calling "get". For example: make get DEP="github.com/zendesk/X@X.X.X")
endif
	go get -mod=readonly $(DEP)
