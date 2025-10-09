# Contributing to Divoom PIXOO64 Golang API

Thank you for your interest in contributing! Here are some guidelines to help you get started.

## How to Contribute

1. **Fork the repository**
2. **Create a feature branch** (`git checkout -b feature/amazing-feature`)
3. **Commit your changes** (`git commit -m 'Add some amazing feature'`)
4. **Push to the branch** (`git push origin feature/amazing-feature`)
5. **Open a Pull Request**

## Development Setup

### Prerequisites

- Go 1.21 or higher
- A Divoom PIXOO64 device for testing (optional)

### Getting Started

```bash
# Clone your fork
git clone https://github.com/eugene-bert/divoom-golang-api.git
cd divoom-golang-api

# Install dependencies
go mod download

# Build the project
go build -v ./...

# Run examples (requires device)
go run examples/main.go
```

## Code Style

- Follow standard Go conventions and idioms
- Use `gofmt` to format your code
- Write clear, descriptive commit messages
- Add comments for exported functions and types
- Update documentation when adding new features

## Testing

While this project currently doesn't have automated tests (contributions welcome!), please test your changes with a real PIXOO64 device if possible.

## Adding New API Methods

When adding new API methods:

1. Add types to `types.go` if needed
2. Implement the method in the appropriate file (`device.go`, `channel.go`, etc.)
3. Update `swagger/openapi.yaml` with the new endpoint
4. Add examples to `examples/main.go`
5. Update `API_REFERENCE.md`

## Questions?

Feel free to open an issue for any questions or concerns!
