# Contributing

## Setup

```bash
go build ./...
go test -race ./...
```

## Stack

- Go library (no main package)
- Zero external dependencies
- HTTP API to Divoom PIXOO64 LED display

## Testing

Tests require a PIXOO64 device on the local network or mock the HTTP layer.

## PR Flow

1. Fork & branch from `main`
2. `go vet ./...` and `go test -race ./...` must pass
3. CI runs automatically on PR
