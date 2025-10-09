# Divoom PIXOO64 Golang API

A Golang API adapter for the Divoom PIXOO64 device.

## Installation

```bash
go get github.com/eugene-bert/divoom-golang-api
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/eugene-bert/divoom-golang-api"
)

func main() {
    // Create a new client with your device's IP address
    client := divoom.NewClient("192.168.1.100")

    // Set brightness
    err := client.SetBrightness(80)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }

    // Send text to display
    err = client.SendText(divoom.TextParams{
        TextID:     1,
        X:          0,
        Y:          0,
        Direction:  0,
        Font:       4,
        TextWidth:  64,
        Speed:      10,
        TextString: "Hello, PIXOO64!",
        Color:      "#00FF00",
        Align:      2, // center
    })
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}
```

## Features

- **Device Management**: Reboot, get/set time, screen control
- **Display Settings**: Brightness, rotation, mirror mode, white balance
- **Channel Settings**: Get all config, set gallery time
- **Drawing**: Send text, images, GIFs, pixel data
- **Tools**: Timer, stopwatch, scoreboard, noise detection
- **System**: Time zone, location, weather info
- **Swagger Documentation**: OpenAPI 3.0 specification with embedded Swagger UI

## Documentation

### Interactive API Documentation

Run the Swagger UI server for interactive documentation:

```bash
go run cmd/swagger-server/main.go
```

Then open your browser: `http://localhost:8080/swagger`

### Embedding Swagger UI in Your Application

```go
package main

import (
    "net/http"
    "github.com/eugene-bert/divoom-golang-api/swagger"
)

func main() {
    http.HandleFunc("/swagger", swagger.Handler())
    http.HandleFunc("/openapi.yaml", swagger.SpecHandler())
    http.ListenAndServe(":8080", nil)
}
```

## API Categories

### Device Management
- `SysReboot()` - Reboot the device
- `GetDeviceTime()` - Get current device time
- `SetUTC(utc int64)` - Set UTC time
- `OnOffScreen(on bool)` - Turn screen on/off

### Display Settings
- `SetBrightness(brightness int)` - Set brightness (0-100)
- `SetScreenRotationAngle(angle int)` - Set rotation (0, 90, 180, 270)
- `SetMirrorMode(enabled bool)` - Enable/disable mirror mode
- `SetWhiteBalance(r, g, b int)` - Set white balance

### Drawing
- `SendText(params TextParams)` - Display text
- `SendGif(params GifParams)` - Display GIF animation
- `ClearText()` - Clear all text areas
- `ResetGifID()` - Reset GIF ID counter

### Tools
- `SetTimer(params TimerParams)` - Set a timer
- `SetStopWatch(params StopWatchParams)` - Set stopwatch
- `SetScoreBoard(params ScoreBoardParams)` - Display scoreboard

## Project Structure

```
divoom-golang-api/
├── client.go              # Core HTTP client
├── types.go               # Type definitions
├── device.go              # Device management
├── channel.go             # Display/channel settings
├── drawing.go             # Drawing commands
├── tools.go               # Tools (timer, stopwatch, etc.)
├── system.go              # System settings
├── swagger/               # Swagger documentation package
│   ├── openapi.yaml       # OpenAPI 3.0 specification
│   ├── swagger.go         # Embedded Swagger UI
│   └── README.md          # Swagger documentation
├── cmd/
│   └── swagger-server/    # Swagger UI server executable
│       └── main.go
├── examples/
│   └── main.go            # Usage examples
├── README.md              # This file
├── API_REFERENCE.md       # Complete API reference
└── GETTING_STARTED.md     # Getting started guide
```

## License

MIT
