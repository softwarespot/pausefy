# Pausefy

Pause Spotify when the `mute` key on the keyboard is pressed for Debian based OS'es

![Demo](pausefy-anime.gif)

## Requirements

-   Debian based OS e.g. Ubuntu
-   Spotify installed from https://www.spotify.com/us/download/linux/ (not the snap version)
-   go 1.23+ - https://golang.org
-   make - https://www.gnu.org/software/make/manual/make.html

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
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.49.0 golangci-lint run -v --tests=false --disable-all -E deadcode,durationcheck,errorlint,exhaustive,gocritic,gosimple,ifshort,ineffassign,misspell,predeclared,revive,staticcheck,unparam,unused,whitespace --max-issues-per-linter=10000 --max-same-issues=10000
```

Local

```bash
GOPATH=$HOME/go golangci-lint run --tests=false --disable-all -E deadcode,durationcheck,errorlint,exhaustive,gocritic,gosimple,ifshort,ineffassign,misspell,predeclared,revive,staticcheck,unparam,unused,whitespace --max-issues-per-linter=10000 --max-same-issues=10000
```
