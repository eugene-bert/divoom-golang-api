# Getting Started with Divoom PIXOO64 Golang API

This guide will help you get up and running with the Divoom PIXOO64 API in minutes.

## Prerequisites

- Go 1.21 or higher installed
- A Divoom PIXOO64 device
- Device connected to your local network
- Device IP address (find it in the Divoom mobile app settings)

## Installation

```bash
go get github.com/eugene-bert/divoom-golang-api
```

## Find Your Device IP

1. Open the Divoom mobile app
2. Connect to your PIXOO64 device
3. Go to Settings → Device Info
4. Note the IP address (e.g., `192.168.1.100`)

## Your First Program

Create a file called `main.go`:

```go
package main

import (
    "log"
    "github.com/eugene-bert/divoom-golang-api"
)

func main() {
    // Replace with your device's IP address
    client := divoom.NewClient("192.168.1.100")

    // Display text - one line!
    if err := client.DisplayText("Hello PIXOO64!", "#00FF00"); err != nil {
        log.Fatal(err)
    }

    // Success! You should see green text on your display
}
```

Run it:

```bash
go run main.go
```

**✅ You should see "Hello PIXOO64!" in green on your display!**

## Next Steps

### 1. Try Different Colors

```go
// Red text
client.DisplayText("Red Text!", "#FF0000")

// Blue text
client.DisplayText("Blue Text!", "#0000FF")

// Yellow text
client.DisplayText("Yellow!", "#FFFF00")

// White text
client.DisplayText("White!", "#FFFFFF")
```

### 2. Custom Position and Font

```go
client.DisplayText("Top Left!", "#FF0000",
    divoom.WithPosition(0, 0),      // x=0, y=0
    divoom.WithFont(2),              // Smaller font
    divoom.WithAlignment(1),         // Left align
)

client.DisplayText("Centered!", "#00FF00",
    divoom.WithPosition(0, 24),      // Middle Y
    divoom.WithFont(4),              // Medium font
    divoom.WithAlignment(2),         // Center align
)
```

### 3. Scrolling Text

```go
client.DisplayText("This text scrolls!", "#0000FF",
    divoom.WithScroll(100, 0),  // speed: 100ms, direction: 0=left
)
```

## Learn More

- [README.md](../README.md) - Full documentation
- [API_REFERENCE.md](API_REFERENCE.md) - Complete API reference
- [TROUBLESHOOTING.md](TROUBLESHOOTING.md) - Common issues and fixes
