# Go Release Action

A comprehensive GitHub action for Go projects that provides linting, testing, building, Docker image creation, versioning, and automatic releases.

## Features

- 🔍 **Linting** with golangci-lint
- 🧪 **Testing** with Go test suite
- 🏗️ **Cross-platform binary builds** for multiple architectures
- 🐳 **Multi-platform Docker images**
- 📦 **Automatic versioning** with svu
- 🚀 **Automatic GitHub releases** with changelog generation
- 🎯 **Matrix builds** for comprehensive testing

## Usage

### Basic Usage

```yaml
name: Release
on:
  push:
    tags:
      - 'v*'

jobs:
  release:
    runs-on: ubuntu-latest
    permissions:
      contents: write
      packages: write
    steps:
      - name: Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Release
        uses: flanksource/go-release-action@v1
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
```

### Advanced Usage

```yaml
name: CI/CD
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  release:
    runs-on: ubuntu-latest
    permissions:
      contents: write
      packages: write
    steps:
      - name: Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Release
        uses: flanksource/go-release-action@v1
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          go-version: '1.21'
          enable-lint: 'true'
          enable-test: 'true'
          enable-build: 'true'
          enable-docker: 'true'
          enable-release: 'true'
          docker-platforms: 'linux/amd64,linux/arm64'
          build-platforms: 'linux/amd64,linux/arm64,darwin/amd64,darwin/arm64,windows/amd64'
          binary-name: 'myapp'
```

## Inputs

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `go-version` | Go version to use | No | `1.21` |
| `github-token` | GitHub token for releases | Yes | - |
| `enable-lint` | Enable linting with golangci-lint | No | `true` |
| `enable-test` | Enable testing | No | `true` |
| `enable-build` | Enable building binaries | No | `true` |
| `enable-docker` | Enable Docker image building | No | `true` |
| `enable-release` | Enable automatic releases | No | `true` |
| `docker-platforms` | Docker platforms to build for | No | `linux/amd64,linux/arm64` |
| `build-platforms` | Platforms to build binaries for | No | `linux/amd64,linux/arm64,darwin/amd64,darwin/arm64,windows/amd64` |
| `binary-name` | Name of the binary to build | No | `app` |

## Outputs

| Output | Description |
|--------|-------------|
| `version` | The version that was released |
| `release-url` | URL of the created release |

## Prerequisites

### Required Files

1. **Makefile** (optional but recommended) - Should include `build` and `test` targets
2. **Dockerfile** (optional) - For Docker image building
3. **go.mod** - Go module file
4. **.golangci.yml** (optional) - golangci-lint configuration

### Example Makefile

```makefile
BINARY_NAME=app
VERSION?=$(shell git describe --tags --always --dirty)
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

build:
	@mkdir -p dist
	go build $(LDFLAGS) -o dist/$(BINARY_NAME) .

test:
	go test -v -race -coverprofile=coverage.out ./...

.PHONY: build test
```

### Example Dockerfile

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

## Project Structure

Your Go project should have the following structure:

```
your-repo/
├── .github/
│   └── workflows/
│       └── release.yml
├── .golangci.yml
├── Dockerfile
├── Makefile
├── go.mod
├── go.sum
├── main.go
└── main_test.go
```

## Permissions

The action requires the following permissions:

```yaml
permissions:
  contents: write    # For creating releases
  packages: write    # For pushing Docker images
```

## Examples

### Simple CLI Tool

```yaml
name: Release CLI
on:
  push:
    tags: ['v*']

jobs:
  release:
    runs-on: ubuntu-latest
    permissions:
      contents: write
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - uses: flanksource/go-release-action@v1
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          enable-docker: 'false'
          binary-name: 'mycli'
```

### Web Service with Docker

```yaml
name: Release Service
on:
  push:
    tags: ['v*']

jobs:
  release:
    runs-on: ubuntu-latest
    permissions:
      contents: write
      packages: write
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - uses: flanksource/go-release-action@v1
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          binary-name: 'myservice'
          docker-platforms: 'linux/amd64,linux/arm64'
```

## License

MIT License - see LICENSE file for details.