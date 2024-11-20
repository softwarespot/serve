# If the "VERSION" environment variable is not set, then use this value instead
VERSION?=1.0.0
TIME=$(shell date +%FT%T%z)
GOVERSION=$(shell go version | awk '{print $$3}' | sed s/go//)

LDFLAGS=-ldflags "\
	-X github.com/softwarespot/serve/internal/version.Version=${VERSION} \
	-X github.com/softwarespot/serve/internal/version.Time=${TIME} \
	-X github.com/softwarespot/serve/internal/version.User=${USER} \
	-X github.com/softwarespot/serve/internal/version.GoVersion=${GOVERSION} \
	-s \
	-w \
"

build:
	@echo building to bin/serve
	@go build $(LDFLAGS) -o ./bin/serve

install:
	@echo copying bin/serve to $(HOME)/bin/serve
	@mv ./bin/serve $(HOME)/bin/serve

test:
	@go test -cover -v ./...

.PHONY: build install test
