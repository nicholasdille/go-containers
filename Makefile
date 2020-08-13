### CUSTOM VARIABLES
#GO_PACKAGE := github.com/my_name/my_repo
#GO_OUTPUT  := my_binary_name

### PREDEFINED VARIABLES ###
BIN_DIR    := bin
REPO       := $(shell git config --get remote.origin.url)
GO_PACKAGE ?= $(REPO:https://%.git=%)
GO_OUTPUT  ?= $(shell basename $(GO_PACKAGE))
SOURCE     := $(shell find . -type f -name \*.go)
PLATFORM   ?= local
M          := $(shell printf "\033[34;1m▶\033[0m")

.PHONY: all
all: $(BIN_DIR)/$(GO_OUTPUT)

.PHONY: check
check: lint test

.PHONY: clean
clean: ; $(info $(M) Cleaning...)
	@rm -rf $(BIN_DIR)

.PHONY: test
test: ; $(info $(M) Running tests...)
	@DOCKER_BUILDKIT=1 docker build . --target unit-test

.PHONY: lint
lint: ; $(info $(M) Checking style...)
	@DOCKER_BUILDKIT=1 docker build . --target lint

go.mod: $(SOURCE) ; $(info $(M) Updating modules...)
	@if test -f "$@"; then \
	    go mod tidy; \
	else \
	    go mod init $(GO_PACKAGE); \
	fi

$(BIN_DIR):
	@mkdir --parents $(BIN_DIR)

.PHONY: $(BIN_DIR)/$(GO_OUTPUT)
$(BIN_DIR)/$(GO_OUTPUT): go.mod $(BIN_DIR) $(SOURCE) test ; $(info $(M) Building...)
	@DOCKER_BUILDKIT=1 docker build . \
	--target bin \
	--output bin/ \
	--build-arg OUTPUT=$(GO_OUTPUT) \
	--platform ${PLATFORM}

%.sha256: % ; $(info $(M) Creating SHA256 for $*...)
	@echo sha256sum $* > $@

%.asc: % ; $(info $(M) Creating signature for $*...)
	@gpg --local-user $$(git config --get user.signingKey) --sign --armor --detach-sig --yes $*
