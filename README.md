# Divoom PIXOO64 Go API

A complete Golang API client for the Divoom PIXOO64 LED display device.

[![Go Report Card](https://goreportcard.com/badge/github.com/eugene-bert/divoom-golang-api)](https://goreportcard.com/report/github.com/eugene-bert/divoom-golang-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Features

- 🎨 **Text Display** - Easy text rendering with colors, fonts, and positioning
- 🖼️ **Canvas Drawing** - Draw pixels, lines, and rectangles
- 📺 **Channel Management** - Switch between device channels
- 🔆 **Display Control** - Brightness, rotation, screen on/off
- ⏱️ **Tools** - Timer, stopwatch, scoreboard, buzzer
- 🌐 **Zero Dependencies** - Uses only Go standard library

## Installation

```bash
go get github.com/eugene-bert/divoom-golang-api
```

## Quick Start

### Find Your Device

```go
devices, _ := divoom.DiscoverDevices()
fmt.Println(devices[0].DeviceIP) // e.g. "192.168.1.100"
```

### Simple Text Display

```go
client := divoom.NewClient("192.168.1.100")
client.DisplayText("Hello World!", "#00FF00")
```

### Display an Image

```go
client.DisplayImageFile("photo.jpg")
client.DisplayImageURL("https://cataas.com/cat?width=200")
```

### Canvas Drawing

```go
package main

import "github.com/eugene-bert/divoom-golang-api"

func main() {
    client := divoom.NewClient("192.168.1.100")

    // Setup
    client.ResetGifID()
    client.SetChannelIndex(3)
    client.SetCustomPageIndex(1)

    // Create canvas
    canvas := divoom.NewCanvas(client)

    // Draw
    canvas.Clear(0, 0, 20)                          // Dark blue background
    canvas.DrawLine(0, 0, 63, 63, 255, 0, 0)        // Red diagonal line
    canvas.FillRectangle(20, 20, 40, 40, 0, 255, 0) // Green square

    // Display
    canvas.Push()
}
```

### Text with Options

```go
client.DisplayText("Custom!", "#FF0000",
    divoom.WithPosition(10, 20),
    divoom.WithFont(5),
    divoom.WithAlignment(2), // center
)
```

## API Overview

### Core Client

- `NewClient(deviceIP string)` - Create client
- `SetTimeout(duration)` - Set HTTP timeout

### Display Control

- `DisplayText(text, color, opts...)` - Easy text display
- `SendText(params)` - Advanced text control
- `ClearText()` - Clear all text

### Canvas Drawing

- `NewCanvas(client)` - Create drawing canvas
- `Clear(r, g, b)` - Fill background
- `SetPixel(x, y, r, g, b)` - Set single pixel
- `DrawLine(x1, y1, x2, y2, r, g, b)` - Draw line
- `FillRectangle(x1, y1, x2, y2, r, g, b)` - Draw rectangle
- `Push()` - Send to display

### Channel Management

- `SetChannelIndex(index)` - Switch channel (0-4)
- `GetChannelIndex()` - Get current channel
- `SetCustomPageIndex(index)` - Select custom page

### Device Settings

- `SetBrightness(0-100)` - Adjust brightness
- `SetScreenRotationAngle(degrees)` - Rotate display
- `OnOffScreen(on bool)` - Power screen on/off

### Tools

- `SetTimer(params)` - Countdown timer
- `SetStopWatch(params)` - Stopwatch
- `PlayBuzzer(params)` - Play sound

## Important Notes

### Text Display Requirements

1. **Must be on Custom channel (3)**: `client.SetChannelIndex(3)`
2. **Must be on CustomPageIndex 1**: `client.SetCustomPageIndex(1)`
3. **Text overlays animations**: Requires background image
4. **Use DisplayText() for simplicity**: Handles all setup automatically

### Canvas Drawing

- Canvas uses single-frame GIF approach (based on Python Pixoo library)
- Auto-resets counter every 32 frames for stability
- Updates show brief loading animation (hardware limitation)
- Recommended update rate: 1-2 seconds between frames

## Documentation

- [API Reference](docs/API_REFERENCE.md) - Complete API documentation
- [Getting Started](docs/GETTING_STARTED.md) - Detailed tutorial

## Project Structure

```
divoom-golang-api/
├── canvas.go          # Canvas drawing API
├── divoom.go          # Main client
├── types.go           # Type definitions
├── options.go         # Text options
├── helpers.go         # Helper methods
├── channel.go         # Channel control
├── device.go          # Device management
├── drawing.go         # Drawing commands
├── tools.go           # Timer, buzzer, etc.
├── system.go          # System settings
├── docs/              # Documentation
└── examples/          # Example programs

```

## Examples

| Example | Description | Usage |
|---------|-------------|-------|
| [basic](examples/basic) | Text display + canvas drawing | `go run examples/basic/main.go <ip>` |
| [image](examples/image) | Display image file or URL | `go run examples/image/main.go <ip> <path-or-url>` |
| [gif](examples/gif) | Play local GIF animation | `go run examples/gif/main.go <ip> <file.gif>` |
| [scrolltext](examples/scrolltext) | Scrolling text with color | `go run examples/scrolltext/main.go <ip> "text" [color]` |
| [clock](examples/clock) | Real-time digital clock | `go run examples/clock/main.go <ip>` |
| [cpu](examples/cpu) | Live CPU usage graph | `cd examples/cpu && go run . <ip>` |

> The `cpu` example has its own `go.mod` (requires `gopsutil`). All others use only the standard library.

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Contributing

Contributions welcome! Please read [CONTRIBUTING.md](docs/CONTRIBUTING.md).

## Credits

- Divoom API: [Official Documentation](https://doc.divoom-gz.com/web/#/12)
- Inspired by: [Python Pixoo Library](https://github.com/SomethingWithComputers/pixoo)

---

**Simple, Clean, Effective** 🎨
