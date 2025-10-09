# Getting Started with Divoom PIXOO64 Golang API

## Prerequisites

- Go 1.21 or higher
- A Divoom PIXOO64 device connected to your local network
- The device's IP address

## Installation

```bash
go get github.com/eugene-bert/divoom-golang-api
```

## Quick Start

### 1. Find Your Device IP Address

**Method 1: Using the Divoom App**
1. Open the Divoom app on your phone
2. Go to device settings
3. Look for network/device info
4. Note the IP address

**Method 2: Check Your Router**
1. Log into your router admin panel
2. Look for connected devices
3. Find "Pixoo64" in the list

### 2. Create Your First Program

Create a file `main.go`:

```go
package main

import (
    "fmt"
    "log"

    divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
    // Replace with your device's IP
    client := divoom.NewClient("192.168.1.100")

    // Display "Hello World!"
    err := client.SendText(divoom.TextParams{
        TextID:     1,
        X:          0,
        Y:          24,
        Direction:  0,
        Font:       4,
        TextWidth:  64,
        Speed:      0,
        TextString: "Hello World!",
        Color:      "#00FF00",
        Align:      2, // center
    })

    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    fmt.Println("Text sent successfully!")
}
```

### 3. Run Your Program

```bash
go run main.go
```

You should see "Hello World!" displayed on your PIXOO64!

## Common Use Cases

### Display Current Time

```go
package main

import (
    "fmt"
    "time"

    divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
    client := divoom.NewClient("192.168.1.100")

    // Enable 24-hour format
    client.SetTime24Flag(true)

    // Set device time to current time
    client.SetUTC(time.Now().Unix())

    fmt.Println("Time synchronized!")
}
```

### Display Dynamic Text (Updates Every Second)

```go
package main

import (
    "fmt"
    "time"

    divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
    client := divoom.NewClient("192.168.1.100")

    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            currentTime := time.Now().Format("15:04:05")

            client.SendText(divoom.TextParams{
                TextID:     1,
                X:          0,
                Y:          24,
                Direction:  0,
                Font:       4,
                TextWidth:  64,
                Speed:      0,
                TextString: currentTime,
                Color:      "#00FF00",
                Align:      2,
            })

            fmt.Printf("Updated: %s\n", currentTime)
        }
    }
}
```

### Control Brightness Based on Time of Day

```go
package main

import (
    "fmt"
    "time"

    divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
    client := divoom.NewClient("192.168.1.100")

    hour := time.Now().Hour()

    var brightness int
    if hour >= 22 || hour < 7 {
        brightness = 20 // Night mode
    } else if hour >= 7 && hour < 9 {
        brightness = 60 // Morning
    } else {
        brightness = 100 // Day
    }

    err := client.SetBrightness(brightness)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    fmt.Printf("Brightness set to %d%%\n", brightness)
}
```

### Display Weather Information

```go
package main

import (
    "fmt"

    divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
    client := divoom.NewClient("192.168.1.100")

    // Set your location (longitude, latitude)
    client.SetLogAndLat("-74.006", "40.7128") // New York

    // Get weather info
    weather, err := client.GetWeatherInfo()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    // Display weather
    weatherText := fmt.Sprintf("%s %d°C", weather.Weather, weather.CurTemp)

    client.SendText(divoom.TextParams{
        TextID:     1,
        X:          0,
        Y:          24,
        Direction:  0,
        Font:       4,
        TextWidth:  64,
        Speed:      0,
        TextString: weatherText,
        Color:      "#FFFF00",
        Align:      2,
    })

    fmt.Println("Weather displayed!")
}
```

## Next Steps

- Check out [API_REFERENCE.md](API_REFERENCE.md) for complete API documentation
- Look at [examples/main.go](examples/main.go) for more examples
- Read the original API documentation in [api_initial_import.md](api_initial_import.md)

## Troubleshooting

### "Connection refused" error
- Verify the device is powered on and connected to your network
- Check that the IP address is correct
- Ensure your computer and PIXOO64 are on the same network

### "Timeout" error
- The device may be busy or unresponsive
- Try increasing the timeout: `client.SetTimeout(30 * time.Second)`
- Restart the device

### Text not displaying correctly
- Check that the font ID is valid (typically 2-7)
- Ensure TextWidth is appropriate for your text length
- Verify the color is in hex format (e.g., "#FF0000")

## Contributing

Feel free to contribute by:
- Adding more API methods from the documentation
- Improving error handling
- Adding tests
- Writing more examples

## Support

For issues and questions:
- Check the API reference documentation
- Review the original API docs in `api_initial_import.md`
- Open an issue on GitHub
