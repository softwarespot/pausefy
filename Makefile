TARGETS_LINUX=pausefy-linux

TARGETS_ALL=$(TARGETS_LINUX)
OS_BUILDS=build-linux

all: $(OS_BUILDS)

# Linux
build-linux: GOOS=linux
build-linux: GOARCH=amd64
build-linux: $(TARGETS_LINUX)

pausefy%:
	@echo building to bin/$@
	@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/$@

clean:
	@cd bin && $(RM) $(TARGETS_ALL)

start:
	@./stop.sh
	@./start.sh

stop:
	@./stop.sh

.PHONY: clean start stop
