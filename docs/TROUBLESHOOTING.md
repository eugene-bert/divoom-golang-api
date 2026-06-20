# Troubleshooting Guide

## Text Not Displaying on PIXOO64

This is the most common issue. If you're sending text commands but nothing appears on the screen (while timer, buzzer, and scoreboard work), follow these steps:

### Solution: Two Critical Steps Required

**CRITICAL**: To display text on PIXOO64, you need BOTH of these steps:

1. **Switch to Custom channel:**
```go
client.SetChannelIndex(3)
```

2. **Send a GIF/animation first (text overlays on animations):**
```go
client.SendBlankScreen()  // Creates a blank black background
// OR
client.SendColorScreen(0x000000)  // Black background
```

**Why?** The PIXOO64 API requires that text is overlaid on top of a GIF/animation. Text cannot be displayed standalone.

**Channel Index Reference:**
- 0 = Faces (clock faces)
- 1 = Cloud Channel
- 2 = Visualizer
- 3 = **Custom** (for your custom text/images)
- 4 = Black screen

### Complete Working Example

```go
package main

import (
    "fmt"
    "log"
    "time"
    divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
    client := divoom.NewClient("192.168.1.100")

    // EASY WAY: Use the convenience method
    client.SetChannelIndex(3)
    time.Sleep(500 * time.Millisecond)

    err := client.DisplayText("Hello!", "#00FF00")
    if err != nil {
        log.Fatalf("Error: %v", err)
    }

    // MANUAL WAY: Do it step by step
    // Step 1: Switch to Custom channel
    client.SetChannelIndex(3)
    time.Sleep(500 * time.Millisecond)

    // Step 2: Send blank screen (REQUIRED!)
    client.SendBlankScreen()
    time.Sleep(200 * time.Millisecond)

    // Step 3: Send text (overlays on blank screen)
    err = client.SendText(divoom.TextParams{
        TextID:     1,
        X:          0,
        Y:          24,
        Direction:  0,
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

    fmt.Println("✓ Text should now be visible!")
}
```

### Quick Test

Run the simple text example:

```bash
go run cmd/simple-text/main.go
```

This example includes all the necessary setup steps.

## Other Common Issues

### Connection Issues

**Problem**: `connection refused` or `timeout` errors

**Solutions**:
1. Verify device IP address is correct
2. Check device is powered on and connected to WiFi
3. Ensure your computer is on the same network
4. Try pinging the device: `ping 192.168.1.100`
5. Increase timeout: `client.SetTimeout(30 * time.Second)`

### Finding Device IP Address

**Method 1: Divoom App**
1. Open Divoom app
2. Tap on your device
3. Go to Settings → Device Info
4. Note the IP address

**Method 2: Router**
1. Log into your router
2. Look for connected devices
3. Find "Pixoo64" or check MAC address starting with "a8:03:2a"

**Method 3: Network Scan**
```bash
# On macOS/Linux
arp -a | grep -i divoom

# Or use nmap
nmap -sn 192.168.1.0/24
```

### Text Appears Garbled or Wrong Font

**Problem**: Text displays but looks wrong

**Solutions**:
1. Try different font IDs (2-7):
   ```go
   Font: 2  // Smallest
   Font: 4  // Medium (recommended)
   Font: 7  // Largest
   ```

2. Adjust Y position for vertical centering:
   ```go
   Y: 24  // Center of 64-pixel screen
   ```

3. Check text width:
   ```go
   TextWidth: 64  // Full screen width
   ```

### Colors Not Working

**Problem**: Text appears but in wrong color

**Solutions**:
1. Verify hex color format:
   ```go
   Color: "#FF0000"  // Red (correct)
   Color: "FF0000"   // Wrong - missing #
   Color: "#F00"     // Wrong - must be 6 digits
   ```

2. Try basic colors:
   - Red: `#FF0000`
   - Green: `#00FF00`
   - Blue: `#0000FF`
   - White: `#FFFFFF`
   - Yellow: `#FFFF00`

### Timer/Scoreboard Works But Text Doesn't

This confirms the issue is **not switching to the Custom channel**.

Tools (timer, stopwatch, scoreboard, buzzer) work on any channel, but custom text/images require switching to channel 3:

```go
client.SetChannelIndex(3)  // Switch to Custom channel
```

### Device Randomly Switches Back to Clock

The PIXOO64 may switch back to clock/gallery channel after some time. You may need to:

1. Call `SetChannelIndex(3)` periodically to keep it on Custom channel
2. Disable auto-rotation in device settings via the app
3. Send updates frequently to keep custom mode active

### Multiple Text Areas

If you're displaying multiple text items simultaneously, use different `TextID` values:

```go
// First text
client.SendText(divoom.TextParams{
    TextID: 1,
    TextString: "Line 1",
    Y: 10,
    // ...
})

// Second text
client.SendText(divoom.TextParams{
    TextID: 2,
    TextString: "Line 2",
    Y: 30,
    // ...
})
```

### API Returns Error Code

Check the error code in the response:

```go
config, err := client.GetAllConf()
if err != nil {
    log.Printf("Error: %v", err)
}
// Also check config.ErrorCode
if config.ErrorCode != 0 {
    log.Printf("Device error code: %d", config.ErrorCode)
}
```

Common error codes:
- `0` - Success
- Non-zero - Device-specific error (check device logs)

## Getting Help

If you're still having issues:

1. Run the full example: `go run examples/main.go`
2. Check what works and what doesn't
3. Open an issue on GitHub with:
   - Your device IP (you can anonymize it)
   - Code you're running
   - Error messages
   - What works vs. what doesn't

## Debug Mode

Add verbose logging to see what's being sent:

```go
package main

import (
    "fmt"
    "log"
    divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
    client := divoom.NewClient("192.168.1.100")

    fmt.Println("1. Switching to Custom channel...")
    if err := client.SetChannelIndex(3); err != nil {
        log.Printf("ERROR: %v", err)
        return
    }
    fmt.Println("   ✓ Switched to Custom channel")

    fmt.Println("2. Sending text...")
    err := client.SendText(divoom.TextParams{
        TextID:     1,
        TextString: "Test",
        X:          0,
        Y:          24,
        Font:       4,
        Color:      "#00FF00",
        TextWidth:  64,
        Align:      2,
    })
    if err != nil {
        log.Printf("ERROR: %v", err)
        return
    }
    fmt.Println("   ✓ Text sent")

    fmt.Println("\n✓ All commands completed successfully!")
    fmt.Println("If text is not visible, check device physically.")
}
```

Run this and verify each step succeeds.
