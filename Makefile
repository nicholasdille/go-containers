### CUSTOM VARIABLES
#GO_PACKAGE := github.com/my_name/my_repo
#GO_OUTPUT  := my_binary_name

### PREDEFINED VARIABLES ###
ROOT_DIR   := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
BIN_DIR    := $(ROOT_DIR)/bin
GOPATH_DIR := $(ROOT_DIR)/.gopath
REPO       := $(shell git config --get remote.origin.url)
GO_PACKAGE ?= $(REPO:https://%.git=%)
GO_OUTPUT  ?= $(shell basename $(GO_PACKAGE))

GIT_COMMIT := $(shell git rev-list -1 HEAD)
BUILD_TIME := $(shell date +%Y%m%d-%H%M%S)
GIT_TAG    := $(shell git describe --tags 2>/dev/null)

M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY:
all: bin

.PHONY:
bin: $(BIN_DIR)/$(GO_OUTPUT)

.PHONY:
clean: $(info $(M) Cleaning...)
	@rm -rf $(GOPATH_DIR)
	@rm -rf $(BIN_DIR)

.PHONY:
test: $(info $(M) Testing...)
	@echo ROOT_DIR=$(ROOT_DIR)
	@echo GO_OUTPUT=$(GO_OUTPUT)

go.mod:
	@go mod init $(GO_PACKAGE)

$(BIN_DIR):
	@mkdir --parents $(BIN_DIR)

$(GOPATH_DIR):
	@mkdir --parents $(GOPATH_DIR)
	@mkdir --parents $(GOPATH_DIR)/src/$(GO_PACKAGE)
	@rmdir $(GOPATH_DIR)/src/$(GO_PACKAGE)
	@ln -sf $(ROOT_DIR) $(GOPATH_DIR)/src/$(GO_PACKAGE)

$(BIN_DIR)/$(GO_OUTPUT): $(GOPATH_DIR) $(BIN_DIR) $(info $(M) Building...)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags "-s -w -X main.GitCommit=$(GIT_COMMIT) -X main.BuildTime=$(BUILD_TIME) -X main.Version=$(GIT_TAG)" -o $@ $(ROOT_DIR)
