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
ROOT_DIR := $(abspath .)
BUILD_ROOT := $(ROOT_DIR)/build
OUT_ROOT := $(ROOT_DIR)/dist
DOCS_ROOT := $(ROOT_DIR)/docs
DPKG_SKEL := $(BUILD_ROOT)/dpkg
BIN_OUT := $(OUT_ROOT)/$(BIN)
PKG_DIR := $(OUT_ROOT)/$(BIN)_$(VERSION)

# Generated asset dirs inside the package skeleton
PKG_BIN := $(PKG_DIR)/usr/local/bin/$(BIN)
PKG_MAN_DIR := $(PKG_DIR)/usr/local/man/man1
PKG_BASH_DIR := $(PKG_DIR)/usr/local/lib/$(BIN)
PKG_ZSH_DIR := $(PKG_DIR)/usr/share/zsh/vendor-completions
PKG_FISH_DIR := $(PKG_DIR)/usr/share/fish/vendor_completions.d

DOCS_FORMAT := markdown

CONTROL_FILE := $(PKG_DIR)/DEBIAN/control

# ---- Go build flags ----
export CGO_ENABLED ?= 0
LDFLAGS := -X $(MODULE_PATH)/pkg/cmd.Version=$(VERSION)

# ---- Helpers ----
define need
	@command -v $(1) >/dev/null 2>&1 || { echo "ERROR: missing dependency: $(1)"; exit 1; }
endef

define ensure_skel
	@test -d "$(DPKG_SKEL)" || { echo "ERROR: missing dpkg skeleton at $(DPKG_SKEL)"; exit 1; }
endef

# ---- Default target ----
.DEFAULT_GOAL := help

.PHONY: help
help: ## Show help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make <target> [VERSION=X]\n\nTargets:\n"} \
	/^[a-zA-Z0-9_.-]+:.*##/ { printf "  %-18s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

# -------- Dev targets --------
.PHONY: build
build: ## Build local dev binary -> ./dist/anetctl
	$(call need,go)
	mkdir -p "$(OUT_ROOT)"
	go build -trimpath -ldflags '$(LDFLAGS)' -o "$(BIN_OUT)" ./

.PHONY: run
run: build ## Run the dev binary (best-effort --version)
	"$(BIN_OUT)" --version || true

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
	rm -rf "$(DOCS_ROOT)/*"

.PHONY: build_docs
build_docs: ## Build docs from the package
	mkdir -p "$(DOCS_ROOT)"
	go run "$(ROOT_DIR)" gen-docs -f "$(DOCS_FORMAT)" "$(DOCS_ROOT)"

# -------- Release pipeline --------
.PHONY: release
release: clean_stage stage_skel compile stage_bin gen_man gen_completions stamp_version build_deb prune_stage ## Full release: dpkg .deb with man/completions

# Step 1: prep stage dir & copy dpkg skeleton
.PHONY: clean_stage
clean_stage: ## Remove existing versioned stage dir
	rm -rf "$(OUT_ROOT)"
	rm -rf "$(DOCS_ROOT)"


.PHONY: stage_skel
stage_skel: ## Copy dpkg skeleton into versioned build dir
	$(call ensure_skel)
	mkdir -p "$(OUT_ROOT)"
	cp -a "$(DPKG_SKEL)" "$(PKG_DIR)"

# Compile into anetctl with ldflags
.PHONY: compile
compile: ## Compile binary with version ldflags
	$(call need,go)
	go build -trimpath -ldflags '$(LDFLAGS)' -o "$(BIN_OUT)" ./

# Place binary into package skeleton
.PHONY: stage_bin
stage_bin: ## Copy binary into package skeleton
	install -Dm755 "$(BIN_OUT)" "$(PKG_BIN)"

# Generate man pages -> gzip -> stage to /usr/local/man/man1
.PHONY: gen_man
gen_man: ## Generate & stage man pages (via gen-docs)
	mkdir -p "$(OUT_ROOT)/mantemp"
	"$(PKG_BIN)" gen-docs -f man "$(OUT_ROOT)/mantemp/"
	# gzip everything produced (extensions may vary depending on generator)
	find "$(OUT_ROOT)/mantemp" -type f -exec gzip -9 {} \;
	mkdir -p "$(PKG_MAN_DIR)"
	# move gzipped man files into man1
	find "$(OUT_ROOT)/mantemp" -maxdepth 1 -type f -name '*.gz' -exec mv {} "$(PKG_MAN_DIR)/" \;

# Generate bash/zsh/fish completions
.PHONY: gen_completions
gen_completions: ## Generate & stage shell completions
	mkdir -p "$(PKG_BASH_DIR)" "$(PKG_ZSH_DIR)" "$(PKG_FISH_DIR)"
	"$(PKG_BIN)" completion bash > "$(PKG_BASH_DIR)/bash_completion"
	"$(PKG_BIN)" completion zsh  > "$(PKG_ZSH_DIR)/_$(BIN)"
	"$(PKG_BIN)" completion fish > "$(PKG_FISH_DIR)/$(BIN).fish"

# Stamp version into DEBIAN/control
.PHONY: stamp_version
stamp_version: ## Replace VERSION_NUMBER token in control with actual version
	$(call need,sed)
	test -f "$(CONTROL_FILE)" || { echo "ERROR: control file missing at $(CONTROL_FILE)"; exit 1; }
	sed -i "s|VERSION_NUMBER|$(VERSION)|" "$(CONTROL_FILE)"

# Build dpkg
.PHONY: build_deb
build_deb: ## Build .deb with dpkg-deb
	$(call need,dpkg-deb)
	dpkg-deb --build "$(PKG_DIR)"
	@echo
	@echo "Built package:"
	@ls -1 "$(OUT_ROOT)"/$(BIN)_$(VERSION).deb 2>/dev/null || true

# Cleanup stage (keep the .deb)
.PHONY: prune_stage
prune_stage: ## Remove staging temp (keeps the .deb)
	rm -rf "$(PKG_DIR)" "$(OUT_ROOT)/mantemp"

# -------- Housekeeping --------
.PHONY: clean
clean: ## Remove all build dists
	rm -rf "$(OUT_ROOT)"
