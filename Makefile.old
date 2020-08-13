### CUSTOM VARIABLES
#GO_PACKAGE := github.com/my_name/my_repo
#GO_OUTPUT  := my_binary_name

### PREDEFINED VARIABLES ###
ROOT_DIR   := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
BIN_DIR    := bin
REPO       := $(shell git config --get remote.origin.url)
GO_PACKAGE ?= $(REPO:https://%.git=%)
GO_OUTPUT  ?= $(shell basename $(GO_PACKAGE))

SOURCE     := $(shell find . -type f -name \*.go)

GIT_COMMIT := $(shell git rev-list -1 HEAD)
BUILD_TIME := $(shell date +%Y%m%d-%H%M%S)
GIT_TAG    := $(shell git describe --tags 2>/dev/null)

M = $(shell printf "\033[34;1m▶\033[0m")

.DEFAULT_GOAL := $(BIN_DIR)/$(GO_OUTPUT)

.PHONY: all clean format lint check deps test

all: $(BIN_DIR)/$(GO_OUTPUT)

clean: ; $(info $(M) Cleaning...)
	@rm -rf $(BIN_DIR)

format: ; $(info $(M) Formatting...)
	@gofmt -l -w $(SOURCE)

lint: ; $(info $(M) Checking style...)
	@if ! type golint >/dev/null 2>&1; then \
	    echo Please install golint: go get -u golang.org/x/lint/golint; \
	fi
	@golint ./...

check: format lint

deps: ; $(info $(M) Generating dependency tree...)
	@if ! type depth >/dev/null 2>&1; then \
	    echo Please install golint: go get github.com/KyleBanks/depth/cmd/depth; \
	fi
	@depth .

test: ; $(info $(M) Running tests...)
	@CGO_ENABLED=0 go test ./...

go.mod: $(SOURCE) ; $(info $(M) Updating modules...)
	@if test -f "$@"; then \
	    go mod tidy; \
	else \
	    go mod init $(GO_PACKAGE); \
	fi

$(BIN_DIR):
	@mkdir --parents $(BIN_DIR)

$(BIN_DIR)/$(GO_OUTPUT): go.mod $(BIN_DIR) $(SOURCE) test ; $(info $(M) Building...)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags "-s -w -X main.GitCommit=$(GIT_COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.Version=$(GIT_TAG)" -o $@ $(ROOT_DIR)

%.sha256: % ; $(info $(M) Creating SHA256 for $*...)
	@echo sha256sum $* > $@

%.asc: % ; $(info $(M) Creating signature for $*...)
	@gpg --local-user $$(git config --get user.signingKey) --sign --armor --detach-sig --yes $*
