# Pausefy

![Go Tests](https://github.com/softwarespot/pausefy/actions/workflows/go.yml/badge.svg)

Pause Spotify when the `mute` key is pressed on Debian-based operating systems.

![Application demo](pausefy-anime.gif)

## Prerequisites

- go 1.26.0 or above
- make (if you want to use the `Makefile` provided)
- Darwin/Debian based OS'es e.g. macOS or Ubuntu

### Darwin based OS'es

- Ensure Spotify is installed from https://www.spotify.com/us/download/mac/
- Requires `osascript` to be installed and running. It is included by default on macOS

### Debian based OS'es

- Ensure Spotify is installed from https://www.spotify.com/us/download/linux/ and **NOT** the snap version, due to using `dbus`
- Requires `dbus` to be installed and running. Install it using your package manager e.g. `sudo apt install dbus`
- Requires `amixer` to be installed, which is part of the `alsa-utils` package. Install it using your package manager e.g. `sudo apt install alsa-utils`
- 

## Installation

Build the binary `pausefy` executable to the directory `./bin` i.e. `./bin/pausefy`.

```bash
make
```

## Usage

### Start and stop

Starts the application `./bin/pausefy` as a background process and stores the process ID (PID) to the file [./scripts/pid](./scripts/pid). The process output i.e. STDOUT is written to the file [./scripts/nohup.out](./scripts/nohup.out).

```bash
make start
```

Stops the application `./bin/pausefy`, using the process ID (PID) stored in the file [./scripts/pid](./scripts/pid) as well as removing the files [./scripts/pid](./scripts/pid) and [./scripts/nohup.out](./scripts/nohup.out).

```bash
make stop
```

### Version

Display the version of the application and exit.

```bash
# As text
./bin/pausefy --version

# As JSON
./bin/pausefy --json --version
```

### Help

Display the help text and exit.

```bash
./bin/pausefy --help
```

## Dependencies

**IMPORTANT:** One dependency is used, which is an adapter for `dbus`.

I could easily use [Cobra](https://github.com/spf13/cobra) (and usually I do,
because it allows me to write powerful CLIs), but I felt it was too much for
such a tiny project. I only ever use dependencies when it's say an adapter for
an external service e.g. Redis, MySQL or Prometheus.

## Linting

Docker

```bash
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run --tests=false --default=none -E durationcheck,errorlint,exhaustive,gocritic,ineffassign,misspell,predeclared,revive,staticcheck,unparam,unused,whitespace --max-issues-per-linter=10000 --max-same-issues=10000
```

Local

```bash
golangci-lint run --tests=false --default=none -E durationcheck,errorlint,exhaustive,gocritic,ineffassign,misspell,predeclared,revive,staticcheck,unparam,unused,whitespace --max-issues-per-linter=10000 --max-same-issues=10000
```

## License

The code has been licensed under the [MIT](https://opensource.org/license/mit) license.
