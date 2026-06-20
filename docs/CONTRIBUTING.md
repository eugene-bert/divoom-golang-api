# Contributing

## How to Contribute

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/my-feature`)
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

## Development Setup

- Go 1.26 or higher
- A Divoom PIXOO64 device for testing (optional — canvas tests run without hardware)

```bash
git clone https://github.com/eugene-bert/divoom-golang-api.git
cd divoom-golang-api
go build ./...
go test -race ./...
```

## Code Style

- `gofmt` all code
- GoDoc comments on all exported symbols
- Zero external dependencies in core library (examples may use deps with separate `go.mod`)

## Adding New API Methods

1. Add types to `types.go` if needed
2. Implement in the appropriate file (`device.go`, `channel.go`, etc.)
3. Add tests
4. Update `docs/API_REFERENCE.md`

## Questions?

Open an issue.
