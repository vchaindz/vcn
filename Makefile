SHELL=/bin/bash -o pipefail

VERSION=0.4.1
TARGETS=linux/amd64 windows/amd64 darwin/amd64

GO ?= go
DOCKER ?= docker

PWD=$(shell pwd)
LDFLAGS := -s -X github.com/vchain-us/vcn/pkg/meta.version=v${VERSION}
TEST_FLAGS ?= -v -race

.PHONY: vcn
vcn:
	$(GO) build ./cmd/vcn

.PHONY: vendor
vendor:
	$(GO) mod vendor

.PHONY: test
test:
	$(GO) vet ./...
	$(GO) test ${TEST_FLAGS} ./...

.PHONY: install
install: TEST_FLAGS=-v
install: vendor test
	$(GO) install -ldflags '${LDFLAGS}' ./cmd/vcn

.PHONY: build/xgo
build/xgo: 
	$(DOCKER) build \
			-f ./build/xgo/Dockerfile \
			-t vcn-xgo \
			./build/xgo

.PHONY: clean/dist
clean/dist: 
	rm -Rf ./dist

.PHONY: clean
clean: clean/dist
	rm -f ./vcn

.PHONY: dist
dist: clean/dist build/xgo
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

.PHONY: dist/NSIS
dist/NSIS:
	mkdir -p dist/NSIS
	cp ./dist/vcn-v${VERSION}-windows-4.0-amd64.exe ./dist/NSIS/vcn.exe
	cp ./resources/NSIS/* ./dist/NSIS/
	$(DOCKER) run --rm \
			-v ${PWD}/dist/NSIS/:/app \
			wheatstalk/makensis:3 /app/setup.nsi
	mv ./dist/NSIS/*_setup.exe ./dist/
	rm -Rf ./dist/NSIS

.PHONY: dist/sign
dist/sign: vcn
	ls ./dist/* | xargs ./vcn sign -y