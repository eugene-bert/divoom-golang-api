# Recent Fixes - Text Display Issue Resolved

## Problem Summary
Text was not displaying on PIXOO64, even though timer, buzzer, and scoreboard commands worked fine.

## Root Cause
I was using the **wrong API command**. The code was calling `UseHTTPCommandSource(true)` which:
1. Is actually for loading commands from a URL file (not what we need)
2. Was causing the device to restart/reconnect
3. Was timing out due to the device reboot

## Solution
Use the correct channel switching command: **`SetChannelIndex(3)`**

### Channel Index Values:
- `0` = Faces (clock faces)
- `1` = Cloud Channel
- `2` = Visualizer
- `3` = **Custom** ← This is what you need for text/images!
- `4` = Black screen

## What Was Fixed

### 1. New API Methods Added (`channel.go`)
```go
client.SetChannelIndex(3)      // Switch to Custom channel
client.GetChannelIndex()        // Get current channel
client.SetCustomPageIndex(0)    // Set custom page (0-2)
```

### 2. Updated All Examples
- `examples/main.go` - Fixed to use `SetChannelIndex(3)`
- `cmd/simple-text/main.go` - Fixed to use `SetChannelIndex(3)`

### 3. Fixed Weather Temperature Type
- Changed `WeatherInfo.CurTemp` from `int` to `float64`
- Fixed printf format from `%d` to `%.1f`

### 4. Updated All Documentation
- `README.md` - Updated quick start
- `GETTING_STARTED.md` - Updated all examples
- `TROUBLESHOOTING.md` - Complete rewrite with correct solution
- `API_REFERENCE.md` - Added new channel methods

## How to Use (Correct Way)

```go
package main

import (
    "fmt"
    "log"
    "time"
    divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
    client := divoom.NewClient("192.168.1.180")

    // Step 1: Switch to Custom channel (CRITICAL!)
    if err := client.SetChannelIndex(3); err != nil {
        log.Fatalf("Failed to switch channel: %v", err)
    }

    // Step 2: Wait for channel switch
    time.Sleep(500 * time.Millisecond)

    // Step 3: Send text
    err := client.SendText(divoom.TextParams{
        TextID:     1,
        X:          0,
        Y:          24,
        Font:       4,
        TextWidth:  64,
        Speed:      0,
        TextString: "Hello PIXOO64!",
        Color:      "#00FF00",
        Align:      2,
    })

    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    fmt.Println("✓ Text displayed!")
}
```

## Quick Test

Run the simple text example:
```bash
go run cmd/simple-text/main.go
```

This should now display text correctly on your PIXOO64 at 192.168.1.180!

## Technical Details

### Why Timer/Buzzer Worked But Text Didn't
- Tools (timer, stopwatch, scoreboard, buzzer) work on **any channel**
- Custom displays (text, images) **only work on Channel 3 (Custom)**
- That's why some commands worked while text commands didn't

### Why the Old Code Failed
- `UseHTTPCommandSource` was triggering a device reboot
- The 10-second default timeout wasn't enough for the reboot
- Even after reboot, the device wasn't in the right channel for custom text

## All Fixed! 🎉

The package now correctly:
- ✅ Switches to the Custom channel
- ✅ Handles weather temperature as float
- ✅ Has complete documentation
- ✅ Includes working examples
- ✅ Passes all builds and vet checks
