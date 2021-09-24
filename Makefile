# # #                                             # # #
# To get started, run make help from the project root #
# # #                                             # # #
.PHONY: build build-alpine clean test help default

BIN_NAME=audmon

# VERSION := $(shell grep "const Version " pkg/version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE_NAME := "lowellmower/audmon"

help:
	@echo 'Management commands for audmon:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'

build:
	go build -o ./bin/ ./cmd/audmon/
