# Divoom PIXOO64 API Reference

## Client Creation

```go
client := divoom.NewClient("192.168.1.100") // Replace with your device IP
client.SetTimeout(15 * time.Second)         // Optional: set custom timeout
```

## Device Management (device.go)

### SysReboot
Reboots the device.
```go
err := client.SysReboot()
```

### GetDeviceTime
Gets the current device time.
```go
deviceTime, err := client.GetDeviceTime()
// Returns: *DeviceTime with UTCTime and LocalTime
```

### SetUTC
Sets the device UTC time.
```go
err := client.SetUTC(time.Now().Unix())
```

### OnOffScreen
Turns the screen on or off.
```go
err := client.OnOffScreen(true)  // Turn on
err := client.OnOffScreen(false) // Turn off
```

### SetDisTempMode
Sets temperature display mode.
```go
err := client.SetDisTempMode(0) // 0 = Celsius, 1 = Fahrenheit
```

### SetScreenRotationAngle
Sets screen rotation angle.
```go
err := client.SetScreenRotationAngle(0) // 0=normal, 1=90°, 2=180°, 3=270°
```

### SetMirrorMode
Enables or disables mirror mode.
```go
err := client.SetMirrorMode(true)  // Enable
err := client.SetMirrorMode(false) // Disable
```

### SetTime24Flag
Sets 24-hour time display.
```go
err := client.SetTime24Flag(true)  // 24-hour mode
err := client.SetTime24Flag(false) // 12-hour mode
```

### SetHighLightMode
Sets high light mode.
```go
err := client.SetHighLightMode(1) // 0 = off, 1 = on
```

### SetWhiteBalance
Sets white balance (RGB values 0-100).
```go
err := client.SetWhiteBalance(100, 100, 100) // r, g, b
```

### PlayTFGif
Plays a GIF from TF card.
```go
err := client.PlayTFGif("animation.gif")
```

### SaveTFGif
Saves current display to TF card as GIF.
```go
err := client.SaveTFGif()
```

### PlayBuzzer
Plays the device buzzer.
```go
err := client.PlayBuzzer(divoom.BuzzerParams{
    ActiveTimeInCycle: 500,  // 500ms on
    OffTimeInCycle:    500,  // 500ms off
    PlayTotalTime:     2000, // 2 seconds total
})
```

## Display & Channel Settings (channel.go)

### SetBrightness
Sets screen brightness (0-100).
```go
err := client.SetBrightness(80)
```

### GetAllConf
Gets all device configuration settings.
```go
config, err := client.GetAllConf()
// Returns: *AllConfigResponse with all settings
```

### SetSubscribeGalleryTime
Sets gallery subscription time in seconds.
```go
err := client.SetSubscribeGalleryTime(60)
```

### SetChannelIndex
Sets the active channel. **CRITICAL for custom displays!**
```go
err := client.SetChannelIndex(3) // 0=Faces, 1=Cloud, 2=Visualizer, 3=Custom, 4=Black
```

### GetChannelIndex
Gets the current active channel index.
```go
channelIndex, err := client.GetChannelIndex()
// Returns: 0-4
```

### SetCustomPageIndex
Sets the custom page index (0-2). Only works when channel is set to Custom (3).
```go
err := client.SetCustomPageIndex(0)
```

## Drawing (drawing.go)

### SendText
Sends text to the display.
```go
err := client.SendText(divoom.TextParams{
    TextID:     1,
    X:          0,
    Y:          24,
    Direction:  0,        // 0=left, 1=right
    Font:       4,        // Font ID
    TextWidth:  64,       // Width in pixels
    Speed:      10,       // Scroll speed in ms
    TextString: "Hello!",
    Color:      "#00FF00", // Hex color
    Align:      2,        // 1=left, 2=center, 3=right
})
```

### ClearText
Clears all text from display.
```go
err := client.ClearText()
```

### SendGif
Sends GIF animation data.
```go
err := client.SendGif(divoom.GifParams{
    PicNum:     10,      // Number of frames
    PicWidth:   64,      // Width (64 for PIXOO64)
    PicOffset:  0,       // Frame offset
    PicID:      1,       // Picture ID
    PicSpeed:   100,     // Animation speed in ms
    PicData:    base64Data, // Base64 encoded image data
})
```

### ResetGifID
Resets the GIF ID counter.
```go
err := client.ResetGifID()
```

### SendRemote
Sends a remote control command.
```go
err := client.SendRemote(keyCode)
```

### UseHTTPCommandSource
Enables/disables HTTP command source.
```go
err := client.UseHTTPCommandSource(true)  // Enable
err := client.UseHTTPCommandSource(false) // Disable
```

## Tools (tools.go)

### SetTimer
Sets a countdown timer.
```go
err := client.SetTimer(divoom.TimerParams{
    Minute: 5,
    Second: 30,
    Status: 1, // 0=stop, 1=start
})
```

### SetStopWatch
Controls the stopwatch.
```go
err := client.SetStopWatch(divoom.StopWatchParams{
    Status: 1, // 0=stop, 1=start, 2=reset
})
```

### SetScoreBoard
Sets the scoreboard display.
```go
err := client.SetScoreBoard(divoom.ScoreBoardParams{
    BlueScore: 5,
    RedScore:  3,
})
```

### SetNoiseStatus
Controls noise detection.
```go
err := client.SetNoiseStatus(1) // 0=off, 1=on
```

## System Settings (system.go)

### SetLogAndLat
Sets location for weather (longitude and latitude).
```go
err := client.SetLogAndLat("30.29", "20.58")
```

### SetTimeZone
Sets the device time zone.
```go
err := client.SetTimeZone("America/New_York")
```

### GetWeatherInfo
Gets current weather information.
```go
weather, err := client.GetWeatherInfo()
// Returns: *WeatherInfo with weather data
```

## Common Types

### TextParams
```go
type TextParams struct {
    TextID     int    // Text ID
    X          int    // X position
    Y          int    // Y position
    Direction  int    // 0=left, 1=right
    Font       int    // Font ID
    TextWidth  int    // Width in pixels
    Speed      int    // Scroll speed in ms
    TextString string // Text to display
    Color      string // Hex color (e.g., "#FF0000")
    Align      int    // 1=left, 2=center, 3=right
}
```

### GifParams
```go
type GifParams struct {
    PicNum     int    // Number of frames
    PicWidth   int    // Width (64 for PIXOO64)
    PicOffset  int    // Frame offset
    PicID      int    // Picture ID
    PicSpeed   int    // Animation speed in ms
    PicData    string // Base64 encoded image data
}
```

### BuzzerParams
```go
type BuzzerParams struct {
    ActiveTimeInCycle int // Active time in ms
    OffTimeInCycle    int // Off time in ms
    PlayTotalTime     int // Total play time in ms
}
```

## Error Handling

All methods return an error that should be checked:

```go
if err := client.SetBrightness(80); err != nil {
    log.Printf("Error: %v", err)
}
```

## Finding Your Device IP

1. Open the Divoom app on your phone
2. Go to device settings
3. Look for "Device Info" or "Network Settings"
4. Note the IP address shown

Or check your router's connected devices list.
