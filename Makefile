# If the "VERSION" environment variable is not set, then use this value instead
VERSION?=1.0.0
TIME=$(shell date +%FT%T%z)
GOVERSION=$(shell go version | awk '{print $$3}' | sed s/go//)

LDFLAGS=-ldflags "\
	-X github.com/softwarespot/pausefy/internal/version.Version=${VERSION} \
	-X github.com/softwarespot/pausefy/internal/version.Time=${TIME} \
	-X github.com/softwarespot/pausefy/internal/version.User=${USER} \
	-X github.com/softwarespot/pausefy/internal/version.GoVersion=${GOVERSION} \
	-s \
	-w \
"

build:
	@echo building to bin/pausefy
	@GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/pausefy

test:
	go test -cover -v ./...

start:
	@./scripts/stop.sh
	@./scripts/start.sh

stop:
	@./scripts/stop.sh

.PHONY: build test start stio
