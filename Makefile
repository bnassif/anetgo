# Makefile for anetctl
# Usage:
#   make build                # local dev build -> ./build/anetctl
#   make release VERSION=1.2-3  # full dpkg build flow (man pages + completions)
#   make clean

SHELL := bash
.ONESHELL:

# ---- Project metadata ----
PKG_NAME := anetctl
MODULE_PATH := github.com/bnassif/anetctl
VERSION ?= $(shell (git describe --tags --abbrev=0 2>/dev/null) || echo 1.0-1)

# ---- Paths ----
ROOT_DIR := $(abspath .)
BUILD_ROOT := $(ROOT_DIR)/build
OUT_ROOT := $(ROOT_DIR)/artifact
DPKG_SKEL := $(BUILD_ROOT)/dpkg
BIN_OUT := $(OUT_ROOT)/$(PKG_NAME)
PKG_DIR := $(OUT_ROOT)/$(PKG_NAME)_$(VERSION)

# Generated asset dirs inside the package skeleton
PKG_BIN := $(PKG_DIR)/usr/local/bin/$(PKG_NAME)
PKG_MAN_DIR := $(PKG_DIR)/usr/local/man/man1
PKG_BASH_DIR := $(PKG_DIR)/usr/local/lib/$(PKG_NAME)
PKG_ZSH_DIR := $(PKG_DIR)/usr/share/zsh/vendor-completions
PKG_FISH_DIR := $(PKG_DIR)/usr/share/fish/vendor_completions.d

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
build: ## Build local dev binary -> ./build/anetctl
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

# -------- Release pipeline (mirrors build_release.sh) --------
.PHONY: release
release: clean_stage stage_skel compile stage_bin gen_man gen_completions stamp_version build_deb prune_stage ## Full release: dpkg .deb with man/completions

# Step 1: prep stage dir & copy dpkg skeleton
.PHONY: clean_stage
clean_stage: ## Remove existing versioned stage dir
	rm -rf "$(PKG_DIR)"

.PHONY: stage_skel
stage_skel: ## Copy dpkg skeleton into versioned build dir
	$(call ensure_skel)
	mkdir -p "$(BUILD_ROOT)"
	cp -a "$(DPKG_SKEL)" "$(PKG_DIR)"

# Step 2/3: compile into ./build/anetctl with ldflags
.PHONY: compile
compile: ## Compile binary with version ldflags
	$(call need,go)
	go build -trimpath -ldflags '$(LDFLAGS)' -o "$(BIN_OUT)" ./

# Step 4: place binary into package skeleton
.PHONY: stage_bin
stage_bin: ## Copy binary into package skeleton
	install -Dm755 "$(BIN_OUT)" "$(PKG_BIN)"

# Step 5: generate man pages -> gzip -> stage to /usr/local/man/man1
.PHONY: gen_man
gen_man: ## Generate & stage man pages (via gen-docs)
	mkdir -p "$(BUILD_ROOT)/mantemp"
	"$(PKG_BIN)" gen-docs -f man "$(BUILD_ROOT)/mantemp/"
	# gzip everything produced (extensions may vary depending on generator)
	find "$(BUILD_ROOT)/mantemp" -type f -exec gzip -9 {} \;
	mkdir -p "$(PKG_MAN_DIR)"
	# move gzipped man files into man1
	find "$(BUILD_ROOT)/mantemp" -maxdepth 1 -type f -name '*.gz' -exec mv {} "$(PKG_MAN_DIR)/" \;

# Step 6: generate bash/zsh/fish completions
.PHONY: gen_completions
gen_completions: ## Generate & stage shell completions
	mkdir -p "$(PKG_BASH_DIR)" "$(PKG_ZSH_DIR)" "$(PKG_FISH_DIR)"
	"$(PKG_BIN)" completion bash > "$(PKG_BASH_DIR)/bash_completion"
	"$(PKG_BIN)" completion zsh  > "$(PKG_ZSH_DIR)/_$(PKG_NAME)"
	"$(PKG_BIN)" completion fish > "$(PKG_FISH_DIR)/$(PKG_NAME).fish"

# Step 7: stamp version into DEBIAN/control
.PHONY: stamp_version
stamp_version: ## Replace VERSION_NUMBER token in control with actual version
	$(call need,sed)
	test -f "$(CONTROL_FILE)" || { echo "ERROR: control file missing at $(CONTROL_FILE)"; exit 1; }
	sed -i "s|VERSION_NUMBER|$(VERSION)|" "$(CONTROL_FILE)"

# Step 8: build dpkg
.PHONY: build_deb
build_deb: ## Build .deb with dpkg-deb
	$(call need,dpkg-deb)
	dpkg-deb --build "$(PKG_DIR)"
	@echo
	@echo "Built package:"
	@ls -1 "$(OUT_ROOT)"/$(PKG_NAME)_$(VERSION).deb 2>/dev/null || true

# Step 9: cleanup stage (keep the .deb)
.PHONY: prune_stage
prune_stage: ## Remove staging temp (keeps the .deb)
	rm -rf "$(PKG_DIR)" "$(OUT_ROOT)/mantemp"

# -------- Housekeeping --------
.PHONY: clean
clean: ## Remove all build artifacts
	rm -rf "$(OUT_ROOT)"
