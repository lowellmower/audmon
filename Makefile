# # #                                             # # #
# To get started, run make help from the project root #
# # #                                             # # #
.PHONY: build build-alpine build-docker

BIN_NAME=audmon

# VERSION := $(shell grep "const Version " pkg/version/version.go | sed -E 's/.*"(.+)"$$/\1/')
VERSION := "0.0.1"
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

build-alpine:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags '-w -linkmode external -extldflags "-static"-o bin/${BIN_NAME} ./cmd/audmon/'
	#go build -o bin/ ./cmd/audmon/

build-docker:
	@echo "building ${BIN_NAME} ${VERSION}"
	docker build -t lowellmower/audmon:local .
	docker run --volume `pwd`/bin:/go/src/github.com/lowellmower/audmon/bin lowellmower/audmon:local

# Local use explicit to dev setup
push:
	scp ./bin/audmon root@lowellmower.com:/usr/local/bin/
	scp ./bin/audmon ubuntu@ec2-18-237-120-181.us-west-2.compute.amazonaws.com:/usr/local/bin
