# Pausefy

![Go Tests](https://github.com/softwarespot/pausefy/actions/workflows/go.yml/badge.svg)

Pause Spotify when the `mute` key on the keyboard is pressed for Debian based OS'es

![Demo](pausefy-anime.gif)

## Prerequisites

-   Debian based OS e.g. Ubuntu
-   go 1.23.0 or above
-   Spotify installed from https://www.spotify.com/us/download/linux/ (**NOT** the snap version)
-   make

## Build

Build the binary to `./bin` as the executable `pausefy-linux`

```bash
make
```

## Start and stop

**NOTE: Ensure the application is built before executing the script**

Starts the application defined in `./bin/pausefy-linux` as a background process and stores the process ID (PID) to the file `pid` (current working directory). The process output i.e. STDOUT is written to the file `nohup.out` in the current working directory

```bash
make start
```

Stops the application defined in `./bin/pausefy-linux`, using the process ID (PID) stored in the file `pid` (current working directory) as well as removing the files `pid` and `nohup.out` in the current working directory

```bash
make stop
```

## Commit (to main branch), build & deploy

**WARNING: This is dangerous to use and should be run with caution**

```bash
git add -u && git ciane && git push -fu && make && make start
```

## Linting

Docker

```bash
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run -v --tests=false --disable-all -E durationcheck,errorlint,exhaustive,gocritic,gosimple,ineffassign,misspell,predeclared,revive,staticcheck,unparam,unused,whitespace --max-issues-per-linter=10000 --max-same-issues=10000
```

Local

```bash
golangci-lint run --tests=false --disable-all -E durationcheck,errorlint,exhaustive,gocritic,gosimple,ineffassign,misspell,predeclared,revive,staticcheck,unparam,unused,whitespace --max-issues-per-linter=10000 --max-same-issues=10000
```
