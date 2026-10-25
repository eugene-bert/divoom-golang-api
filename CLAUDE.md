# divoom-golang-api

Go API client library for the Divoom PIXOO64 LED display. Supports drawing, animations, text, system info display.

## Commands

```bash
go build ./...          # build
go vet ./...            # lint
go test -race ./...     # test
```

## CI/CD

- **CI** (`go.yml`): `go build`, `go vet`, `go test` on push/PR to main
- **Dependabot**: auto-merge patch/minor updates
- No release CI — this is a Go library, imported via `go get`

## Usage

```go
import "github.com/eugene-bert/divoom-golang-api"
```

## Key files

- Root package — API client, all HTTP communication with PIXOO64
- Sends commands via HTTP POST to device IP on local network
- No external dependencies beyond stdlib

## Important

- This is a **library**, not a standalone service
- Device communicates via HTTP (not MQTT/BLE) on local network
- Device IP must be known (found via Divoom app or network scan)
