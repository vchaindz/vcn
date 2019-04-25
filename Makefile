SHELL=/bin/bash -o pipefail

VERSION=0.4.0
TARGETS=linux/amd64 windows/amd64 darwin/amd64

GO ?= go
DOCKER ?= docker

PWD=$(shell pwd)
LDFLAGS := -s -X github.com/vchain-us/vcn/pkg/meta.version=v${VERSION}

.PHONY: vcn
vcn:
	$(GO) build ./cmd/vcn

.PHONY: vendor
vendor:
	$(GO) mod vendor

.PHONY: test
test:
	$(GO) vet ./...
	$(GO) test -v -race ./...

.PHONY: install
install: vendor test
	$(GO) install -ldflags '${LDFLAGS}' ./cmd/vcn

.PHONY: builder
builder: 
	$(DOCKER) build \
			-f ./build/Dockerfile.builder \
			-t vcn-xgo \
			./build

.PHONY: clean
clean: 
	rm -f ./vcn
	rm -Rf ./dist

.PHONY: dist
dist: clean builder
	mkdir -p dist
	$(DOCKER) run --rm \
			-v ${PWD}/dist:/dist \
			-v ${PWD}:/source:ro \
			-e GO111MODULE=on \
			-e FLAG_LDFLAGS="${LDFLAGS}" \
			-e TARGETS="${TARGETS}" \
			-e PACK=cmd/vcn \
			-e OUT=vcn-v${VERSION} \
			vcn-xgo .