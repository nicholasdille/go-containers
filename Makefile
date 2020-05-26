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

.PHONY:
all: $(BIN_DIR)/$(GO_OUTPUT)

.PHONY:
clean:
	@rm -rf $(GOPATH_DIR)
	@rm -rf $(BIN_DIR)

.PHONY:
test:
	@echo ROOT_DIR=$(ROOT_DIR)
	@echo GO_OUTPUT=$(GO_OUTPUT)

$(BIN_DIR):
	@mkdir --parents $(BIN_DIR)

$(GOPATH_DIR):
	@mkdir --parents $(GOPATH_DIR)
	@mkdir --parents $(GOPATH_DIR)/src/$(GO_PACKAGE)
	@rmdir $(GOPATH_DIR)/src/$(GO_PACKAGE)
	@ln -sf $(ROOT_DIR) $(GOPATH_DIR)/src/$(GO_PACKAGE)

$(BIN_DIR)/$(GO_OUTPUT): $(GOPATH_DIR) $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(GO_OUTPUT) main.go
