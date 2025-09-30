# Example workflows for using go-release-action

## Simple Release Workflow

Create `.github/workflows/release.yml`:

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

## CI Workflow with Pull Request Testing

Create `.github/workflows/ci.yml`:

```yaml
name: CI

on:
  push:
    branches: [main, master]
  pull_request:
    branches: [main, master]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Test
        uses: flanksource/go-release-action@v1
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          enable-release: 'false'
          enable-docker: 'false'
```

## Advanced Configuration

```yaml
name: Advanced Release

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
          go-version: '1.21'
          binary-name: 'myapp'
          docker-platforms: 'linux/amd64,linux/arm64,linux/arm/v7'
          build-platforms: 'linux/amd64,linux/arm64,darwin/amd64,darwin/arm64,windows/amd64,windows/arm64'
```