# Swagger Documentation Package

This package contains OpenAPI/Swagger documentation for the Divoom PIXOO64 API.

## Usage

### Embedded Swagger UI

```go
// Run the ready-made server
go run cmd/swagger-server/main.go

// Or embed in your application:
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/eugene-bert/divoom-golang-api/swagger"
)

func main() {
    http.HandleFunc("/swagger", swagger.Handler())
    http.HandleFunc("/openapi.yaml", swagger.SpecHandler())

    fmt.Println("Swagger UI: http://localhost:8080/swagger")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### Accessing OpenAPI Specification

```go
package main

import (
    "fmt"
    "github.com/eugene-bert/divoom-golang-api/swagger"
)

func main() {
    // Get OpenAPI specification as a string
    spec := swagger.OpenAPISpec
    fmt.Println(spec)
}
```

## Files

- `openapi.yaml` - OpenAPI 3.0 API specification
- `swagger.go` - Go package with embedded specification and Swagger UI
- `README.md` - This documentation

## Viewing Documentation

After starting the server, open your browser and navigate to:
- `http://localhost:8080/swagger` - Interactive Swagger UI
- `http://localhost:8080/openapi.yaml` - Raw YAML specification

## Swagger UI Features

- View all available API endpoints
- Interactive API testing
- View request and response schemas
- Usage examples
- Complete API documentation

## Updating Documentation

To update the documentation, edit the `openapi.yaml` file and rebuild the package:

```bash
go build ./swagger
```
