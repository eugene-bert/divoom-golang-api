# Divoom PIXOO64 Golang API

A complete Golang API client for the Divoom PIXOO64 LED display device.

[![Go Report Card](https://goreportcard.com/badge/github.com/eugene-bert/divoom-golang-api)](https://goreportcard.com/report/github.com/eugene-bert/divoom-golang-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Features

- 🎨 **Display Text** - Easy text display with customizable colors, positions, and fonts
- 🖼️ **Custom Animations** - Upload and display custom GIF animations
- 🌐 **URL GIFs** - Play GIF animations from URLs
- ⏱️ **Tools** - Timer, stopwatch, scoreboard, buzzer
- 🔆 **Display Control** - Brightness, rotation, mirror mode, white balance
- 📺 **Channel Management** - Switch between channels (Faces, Cloud, Visualizer, Custom)
- 📖 **Comprehensive Documentation** - Interactive Swagger UI and detailed examples

## Installation

```bash
go get github.com/eugene-bert/divoom-golang-api
```

## Quick Start

### Simple Text Display (Recommended)

```go
package main

import (
    "log"
    "github.com/eugene-bert/divoom-golang-api"
)

func main() {
    // Create client with your device's IP address
    client := divoom.NewClient("192.168.1.180")

    // Display text - one line!
    if err := client.DisplayText("Hello PIXOO64!", "#00FF00"); err != nil {
        log.Fatal(err)
    }
}
```

### Text with Custom Options

```go
// Custom position, font, and alignment
client.DisplayText("Custom Text!", "#FF0000",
    divoom.WithPosition(10, 20),
    divoom.WithFont(5),
    divoom.WithAlignment(1), // 1=left, 2=center, 3=right
)

// Scrolling text
client.DisplayText("Scrolling!", "#0000FF",
    divoom.WithScroll(100, 0), // speed: 100ms, direction: 0=left
)
```

### Custom Colored Background with Text

```go
// Reset GIF ID before manual operations
client.ResetGifID()
time.Sleep(500 * time.Millisecond)

// Switch to Custom channel page 1
client.SetChannelIndex(3)
time.Sleep(500 * time.Millisecond)
client.SetCustomPageIndex(1)
time.Sleep(2 * time.Second)

// Send blue background
client.SendColorScreen(0x0000FF)
time.Sleep(500 * time.Millisecond)

// Overlay white text
client.SendText(divoom.TextParams{
    TextID:     1,
    X:          0,
    Y:          24,
    Font:       4,
    TextWidth:  64,
    TextString: "Blue Background!",
    Color:      "#FFFFFF",
    Align:      2,
})
```

### Play GIF from URL with Text Overlay

```go
client.SetChannelIndex(3)
time.Sleep(500 * time.Millisecond)
client.SetCustomPageIndex(1)
time.Sleep(2 * time.Second)

// Play GIF from URL
client.PlayTFGif("http://f.divoom-gz.com/64_64.gif")
time.Sleep(1 * time.Second)

// Add text overlay
client.SendText(divoom.TextParams{
    TextID:     1,
    X:          0,
    Y:          24,
    Font:       4,
    TextWidth:  64,
    TextString: "GIF + Text!",
    Color:      "#FFFF00",
    Align:      2,
})
```

## API Overview

### Device Management
- `NewClient(deviceIP string)` - Create a new client
- `SysReboot()` - Reboot the device
- `GetDeviceTime()` - Get current device time
- `SetUTC(utc int64)` - Set UTC time
- `OnOffScreen(on bool)` - Turn screen on/off

### Display Settings
- `SetBrightness(brightness int)` - Set brightness (0-100)
- `SetScreenRotationAngle(angle int)` - Set rotation (0, 90, 180, 270)
- `SetMirrorMode(enabled bool)` - Enable/disable mirror mode
- `SetWhiteBalance(r, g, b int)` - Set white balance (0-100 each)

### Channel Management
- `SetChannelIndex(index int)` - Switch channel (0=Faces, 1=Cloud, 2=Visualizer, 3=Custom, 4=Black)
- `GetChannelIndex()` - Get current channel
- `SetCustomPageIndex(index int)` - Select custom page (0=Favorites, 1=Custom content, 2=Unknown)

### Text Display
- `DisplayText(text, color string, opts...)` - **Easy way** - handles everything automatically
- `SendText(params TextParams)` - Manual text overlay
- `ClearText()` - Clear all text

### GIF/Animation Display
- `SendBlankScreen()` - Send blank black background
- `SendColorScreen(rgb uint32)` - Send solid color background
- `SendGif(params GifParams)` - Upload custom GIF animation
- `PlayTFGif(url string)` - Play GIF from URL
- `ResetGifID()` - **Important**: Reset GIF ID counter before uploading new GIFs

### Tools
- `SetTimer(params TimerParams)` - Set countdown timer
- `SetStopWatch(params StopWatchParams)` - Control stopwatch
- `SetScoreBoard(params ScoreBoardParams)` - Display scoreboard
- `PlayBuzzer(params BuzzerParams)` - Play buzzer sound

## Important Notes

### Text Display Requirements

Text display on PIXOO64 has specific requirements:

1. **Must be on Custom channel (3)** - Use `SetChannelIndex(3)`
2. **Must be on CustomPageIndex 1** - Use `SetCustomPageIndex(1)` (page 0 shows saved favorites)
3. **Text overlays on animations** - Text cannot display standalone; it must overlay a GIF/animation
4. **Use `DisplayText()` for simplicity** - It handles all setup automatically

### GIF Upload Requirements

When uploading custom GIFs:

1. **Call `ResetGifID()` first** - Clears accumulated GIF data
2. **Use 2-frame animations** - Single-frame GIFs may not display
3. **Frames must be different** - Device requires different frame data
4. **Use proper timing delays** - Allow device time to process (500ms-2s between steps)

### Example: Why DisplayText() is Recommended

**Easy way (recommended):**
```go
client.DisplayText("Hello!", "#00FF00")
```

**Manual way (more control, more complex):**
```go
client.ResetGifID()
time.Sleep(500 * time.Millisecond)
client.SetChannelIndex(3)
time.Sleep(500 * time.Millisecond)
client.SetCustomPageIndex(1)
time.Sleep(2 * time.Second)
client.SendBlankScreen()
time.Sleep(500 * time.Millisecond)
client.SendText(divoom.TextParams{
    TextID: 1, X: 0, Y: 24, Font: 4,
    TextWidth: 64, TextString: "Hello!",
    Color: "#00FF00", Align: 2,
})
```

## Documentation

### Interactive API Documentation

Run the built-in Swagger UI server:

```bash
go run cmd/swagger-server/main.go
```

Then open: `http://localhost:8080/swagger`

### Embed Swagger UI in Your Application

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

## Examples

### Run Complete Examples

```bash
# All 4 examples: text, custom position, colored background, GIF+text
go run cmd/working-example/main.go

# Simple one-line text test
go run cmd/final-test/main.go
```

### Create Custom Animation

```go
package main

import (
    "encoding/base64"
    "github.com/eugene-bert/divoom-golang-api"
    "time"
)

func main() {
    client := divoom.NewClient("192.168.1.180")

    // Reset and setup
    client.ResetGifID()
    time.Sleep(500 * time.Millisecond)
    client.SetChannelIndex(3)
    time.Sleep(500 * time.Millisecond)
    client.SetCustomPageIndex(1)
    time.Sleep(2 * time.Second)

    // Create 2-frame RED/BLUE animation
    // Frame 1: RED
    redPixels := make([]byte, 64*64*3)
    for i := 0; i < 64*64; i++ {
        redPixels[i*3] = 255 // R channel
    }
    redData := base64.StdEncoding.EncodeToString(redPixels)

    client.SendGif(divoom.GifParams{
        PicNum: 2, PicWidth: 64, PicOffset: 0,
        PicID: 1, PicSpeed: 500, PicData: redData,
    })
    time.Sleep(200 * time.Millisecond)

    // Frame 2: BLUE
    bluePixels := make([]byte, 64*64*3)
    for i := 0; i < 64*64; i++ {
        bluePixels[i*3+2] = 255 // B channel
    }
    blueData := base64.StdEncoding.EncodeToString(bluePixels)

    client.SendGif(divoom.GifParams{
        PicNum: 2, PicWidth: 64, PicOffset: 1,
        PicID: 1, PicSpeed: 500, PicData: blueData,
    })

    // Overlay text
    time.Sleep(500 * time.Millisecond)
    client.SendText(divoom.TextParams{
        TextID: 1, X: 0, Y: 24, Font: 4,
        TextWidth: 64, TextString: "Flashing!",
        Color: "#FFFFFF", Align: 2,
    })
}
```

## Project Structure

```
divoom-golang-api/
├── client.go              # HTTP client core
├── types.go               # Type definitions
├── device.go              # Device management
├── channel.go             # Channel/display settings
├── drawing.go             # Drawing commands (text, GIF)
├── tools.go               # Tools (timer, stopwatch, etc.)
├── system.go              # System settings
├── helpers.go             # Helper methods (DisplayText, etc.)
├── itemlist.go            # ItemList API
├── swagger/               # Swagger documentation
│   ├── openapi.yaml       # OpenAPI 3.0 spec
│   ├── swagger.go         # Embedded Swagger UI
│   └── README.md
├── cmd/
│   ├── swagger-server/    # Swagger UI server
│   ├── working-example/   # Complete examples
│   ├── final-test/        # Simple test
│   └── [other test commands]
├── docs/                  # Documentation
│   ├── API_REFERENCE.md   # Complete API reference
│   ├── CHANGELOG.md       # Version history
│   ├── CONTRIBUTING.md    # Contribution guidelines
│   ├── GETTING_STARTED.md # Beginner's guide
│   ├── SOLUTION.md        # Technical details
│   └── TROUBLESHOOTING.md # Common issues
├── README.md              # This file
└── LICENSE                # MIT License
```

## Troubleshooting

### Text Not Displaying

1. Ensure you're on **Custom channel (3)** and **CustomPageIndex 1**
2. Use `ResetGifID()` before sending new GIFs
3. Use `DisplayText()` helper - it handles everything automatically

### Loading/Hourglass Animation

This is normal when:
- Switching channels
- Uploading GIF frames
- Processing animations

The timing delays in the code are necessary for reliable operation.

### GIF Not Displaying After SendHttpGif

1. Call `ResetGifID()` first
2. Switch to CustomPageIndex 1 **before** sending GIF
3. Ensure 2 frames with **different** data (not identical)
4. Add proper timing delays between operations

## API Reference

See [API_REFERENCE.md](docs/API_REFERENCE.md) for complete API documentation.

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Contributing

Contributions welcome! Please read [CONTRIBUTING.md](docs/CONTRIBUTING.md) for guidelines.

## Credits

- API Documentation: [Divoom API](https://doc.divoom-gz.com/web/#/12)
- Device: [Divoom PIXOO64](https://divoom.com/)

## Changelog

See [CHANGELOG.md](docs/CHANGELOG.md) for version history.

## Support

- Issues: [GitHub Issues](https://github.com/eugene-bert/divoom-golang-api/issues)
- Documentation: [Swagger UI](http://localhost:8080/swagger) (after running `go run cmd/swagger-server/main.go`)

---

**Made with ❤️ for the PIXOO64 community**
