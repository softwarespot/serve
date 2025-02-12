# Serve

![Go Tests](https://github.com/softwarespot/serve/actions/workflows/go.yml/badge.svg)

This application is a simple static file server that serves files from either the current working directory or specified directory over HTTP with CORS support, allowing easy access for development and testing.

## Prerequisites

- go 1.24.0 or above
- make (if you want to use the `Makefile` provided)

## Installation

### Binaries

`serve` is available on Linux, macOS and Windows platforms.
Binaries for Linux, Windows and Mac are available as tarballs in the [release page](https://github.com/softwarespot/serve/releases).

### Go Install

Install to the Go `bin` directory e.g. `$HOME/go/bin/`.

```bash
go install github.com/softwarespot/serve@latest
```

## Usage

### Serve

Serve files from the current working directory on port `8100`.

```bash
./bin/serve
```

Serve files from the current working directory on port `3000`.

```bash
./bin/serve -p 3000
```

### Version

Display the version of the application and exit.

```bash
# As text
./bin/serve --version

# As JSON
./bin/serve --json --version
```

### Help

Display the help text and exit.

```bash
./bin/serve --help
```

## Dependencies

**IMPORTANT:** No 3rd party dependencies are used.

I could easily use [Cobra](https://github.com/spf13/cobra) (and usually I do,
because it allows me to write powerful CLIs), but I felt it was too much for
such a tiny project. I only ever use dependencies when it's say an adapter for
an external service e.g. Redis, MySQL or Prometheus.

## Linting

Docker

```bash
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run -v --tests=false --disable-all -E durationcheck,errorlint,exhaustive,gocritic,gosimple,ineffassign,misspell,predeclared,revive,staticcheck,unparam,unused,whitespace --max-issues-per-linter=10000 --max-same-issues=10000
```

Local

```bash
golangci-lint run --tests=false --disable-all -E durationcheck,errorlint,exhaustive,gocritic,gosimple,ineffassign,misspell,predeclared,revive,staticcheck,unparam,unused,whitespace --max-issues-per-linter=10000 --max-same-issues=10000
```

## License

The code has been licensed under the [MIT](https://opensource.org/license/mit) license.
