# Copilot Instructions

Go library for Divoom PIXOO64 LED display. HTTP API client.

## Build & Test

```bash
go build ./...
go vet ./...
go test -race ./...
```

## Architecture

- Single Go package, no main — imported as library
- HTTP POST commands to device IP on local network
- No external dependencies
