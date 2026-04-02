# Makefile for anetctl
# Usage:
#   make build                # local dev build -> ./build/anetctl
#   make release VERSION=1.2-3  # full dpkg build flow (man pages + completions)
#   make clean

SHELL := bash
.ONESHELL:

# ---- Project metadata ----
BIN := anetctl
MODULE_PATH := github.com/bnassif/anetgo
VERSION ?= $(shell (git describe --tags --abbrev=0 2>/dev/null) || echo 1.0-1)

# ---- Paths ----
ROOT_DIR := .
BUILD_ROOT := $(ROOT_DIR)/build
DIST_ROOT := $(ROOT_DIR)/dist
DOCS_ROOT := $(ROOT_DIR)/docs

## -- Build Paths --
MAN_DIR := $(BUILD_ROOT)/manpages/man1
MAN_TEMP_DIR := $(BUILD_ROOT)/mantemp
COMPLETION_BASH_DIR := $(BUILD_ROOT)/completions/bash
COMPLETION_FISH_DIR := $(BUILD_ROOT)/completions/fish
COMPLETION_ZSH_DIR := $(BUILD_ROOT)/completions/zsh
BIN_OUT := $(BUILD_ROOT)/bin/$(BIN)

## -- Stamp Paths --
DOCS_STAMP := $(DOCS_ROOT)/$(BIN).md
MAN_STAMP := $(BUILD_ROOT)/.man-stamp
COMP_STAMP := $(BUILD_ROOT)/.comp-stamp

# Add variables for nfpm usage
##############################!!!!!!!!!!!!!!!!!!!!!!!
BUILD_ARCH ?= amd64
BUILD_OS ?= linux
PACKAGER ?= deb

PGP_PRIVATE_KEY_FILE ?= ""
PGP_PRIVATE_KEY_ID ?= 6E0E33C781685C9E
PGP_PRIVATE_KEY_PASSPHRASE ?= ""
RSA_PRIVATE_KEY_FILE ?= ""
RSA_PRIVATE_KEY_PASSPHRASE ?= ""

DOCS_FORMAT := markdown


# ---- Go build flags ----
export CGO_ENABLED ?= 0
LDFLAGS := -X $(MODULE_PATH)/pkg.Version=$(VERSION)

# ---- Helpers ----
define need
	@command -v $(1) >/dev/null 2>&1 || { echo "ERROR: missing dependency: $(1)"; exit 1; }
endef

# ---- Default target ----
.DEFAULT_GOAL := help

.PHONY: help
help: ## Show help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make <target> [VERSION=X]\n\nTargets:\n"} \
	/^[a-zA-Z0-9_.-]+:.*##/ { printf "  %-18s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

# -------- Dev targets --------
.PHONY: run
run: ## Run the dev binary (best-effort version)
	go run "$(ROOT_DIR)" version || true

.PHONY: fmt
fmt: ## go fmt
	go fmt ./...

.PHONY: vet
vet: ## go vet
	go vet ./...

.PHONY: test
test: ## go test
	go test ./...

.PHONY: check
check: fmt vet test ## Run fmt, vet, test

# -------- Docs pipeline --------
.PHONY: docs
docs: clean_docs build_docs

.PHONY: clean_docs
clean_docs: ## Remove existing docs from working tree
	find "$(DOCS_ROOT)" -mindepth 1 -delete
	rm -f "$(DOCS_STAMP)"


$(DOCS_STAMP): ## Build docs from the package
	mkdir -p "$(DOCS_ROOT)"
	go run "$(ROOT_DIR)" gen-docs -f "$(DOCS_FORMAT)" "$(DOCS_ROOT)"
	touch $@

.PHONY: build_docs
build_docs: $(DOCS_STAMP)

# -------- Release pipeline --------
.PHONY: release
release: clean package_all clean_artifacts ## Full release

.PHONY: clean
clean: clean_dist clean_artifacts clean_stamps ## Clean all build artifacts

.PHONY: clean_dist
clean_dist: ## Remove all build dists
	rm -rf "$(DIST_ROOT)"

.PHONY: clean_artifacts
clean_artifacts: ## Remove all build artifacts
	rm -f "$(BIN_OUT)"
	rm -rf "$(MAN_DIR)" "$(MAN_TEMP_DIR)" \
		"$(COMPLETION_BASH_DIR)" "$(COMPLETION_FISH_DIR)" "$(COMPLETION_ZSH_DIR)"

.PHONY: clean_stamps
clean_stamps: ## Remove all build stamps
	rm -f "$(DOCS_STAMP)" "$(MAN_STAMP)" "$(COMP_STAMP)"

$(BIN_OUT): ## Compile binary with version ldflags
	$(call need,go)
	GOOS=$(BUILD_OS) GOARCH=$(BUILD_ARCH) go build -trimpath -ldflags '$(LDFLAGS)' -o "$(BIN_OUT)" ./

.PHONY: bin
bin: $(BIN_OUT)

$(MAN_STAMP): $(BIN_OUT) ## Generate & stage man pages (via gen-docs)
	mkdir -p "$(MAN_TEMP_DIR)" "$(MAN_DIR)"
	go run "$(ROOT_DIR)" gen-docs -f man "$(MAN_TEMP_DIR)"
	cp -r "$(MAN_TEMP_DIR)/" "$(MAN_DIR)/"
	find "$(MAN_DIR)" -type f -name '*.1' -exec gzip -9 {} \;
	touch $@

.PHONY: gen_man
gen_man: $(MAN_STAMP)

$(COMP_STAMP): $(BIN_OUT) ## Generate & stage shell completions
	mkdir -p "$(COMPLETION_BASH_DIR)" "$(COMPLETION_FISH_DIR)" "$(COMPLETION_ZSH_DIR)"
	go run "$(ROOT_DIR)" completion bash > "$(COMPLETION_BASH_DIR)/$(BIN)"
	go run "$(ROOT_DIR)" completion fish > "$(COMPLETION_FISH_DIR)/$(BIN).fish"
	go run "$(ROOT_DIR)" completion zsh  > "$(COMPLETION_ZSH_DIR)/_$(BIN)"
	touch $@

.PHONY: gen_completions
gen_completions: $(COMP_STAMP)

.PHONY: package
package: gen_man gen_completions ## Package the application for a package manager
	$(call need,nfpm)
	mkdir -p "$(DIST_ROOT)"
	BUILDARCH=$(BUILD_ARCH) BUILDOS=$(BUILD_OS) SEMVER=$(VERSION) \
		PGP_PRIVATE_KEY_FILE=$(PGP_PRIVATE_KEY_FILE) PGP_PRIVATE_KEY_ID=$(PGP_PRIVATE_KEY_ID) RSA_PRIVATE_KEY_FILE=$(RSA_PRIVATE_KEY_FILE) \
		NFPM_RPM_PASSPHRASE=$(PGP_PRIVATE_KEY_PASSPHRASE) NFPM_DEB_PASSPHRASE=$(PGP_PRIVATE_KEY_PASSPHRASE) NFPM_APK_PASSPHRASE=$(RSA_PRIVATE_KEY_PASSPHRASE) \
		nfpm package --packager $(PACKAGER) --target $(DIST_ROOT)/

.PHONY: package_deb
package_deb:
	$(MAKE) package PACKAGER=deb

.PHONY: package_rpm
package_rpm:
	$(MAKE) package PACKAGER=rpm

.PHONY: package_apk
package_apk:
	$(MAKE) package PACKAGER=apk

.PHONY: package_all
package_all: package_deb package_rpm