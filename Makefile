# Makefile
PROJECT_ROOT := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

all: build
build:
	@go build -o $(PROJECT_ROOT)/call
	@echo "\033[32m Done. \033[37m"

.PHONY: all build
