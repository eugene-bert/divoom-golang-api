# THE SOLUTION - Text Display on PIXOO64

## TL;DR - Quick Fix

To display text on PIXOO64, you **MUST**:

1. **Reset GIF ID** - `client.ResetGifID()`
2. Switch to **Custom channel (3)** - `client.SetChannelIndex(3)`
3. Switch to **CustomPageIndex 1** (NOT 0!) - `client.SetCustomPageIndex(1)`
4. Send a **2-frame GIF/animation** with **different** frames
5. Then send text overlay

**OR** just use: `client.DisplayText("Hello!", "#00FF00")` - it does everything!

## The Easy Way (Recommended)

```go
client := divoom.NewClient("192.168.1.180")

// One line - handles everything automatically!
client.DisplayText("Hello World!", "#00FF00")
```

## The Problem We Discovered

After extensive testing, we found that:

1. **Text CANNOT display standalone** - it requires a base GIF/animation
2. **Custom channel has 3 pages (0, 1, 2)**:
   - Page 0 = Your saved favorites collection
   - Page 1 = Custom content (what we need!)
   - Page 2 = Unknown (shows test pattern)
3. **SendHttpGif uploads but doesn't auto-display** - you must be on the correct page first
4. **You must be on CustomPageIndex 1 BEFORE sending the GIF** for it to display
5. **ResetGifID() is CRITICAL** - accumulated GIF data from multiple uploads causes loading/hourglass animations
6. **Frames must be DIFFERENT** - device won't display animation with identical frame data
7. **2-frame minimum** - single-frame GIFs don't display reliably

## The Complete Solution

### Method 1: Using the Helper (Easiest)

```go
package main

import (
    "github.com/eugene-bert/divoom-golang-api"
)

func main() {
    client := divoom.NewClient("192.168.1.180")

    // Simple text display
    client.DisplayText("Hello!", "#00FF00")

    // With options
    client.DisplayText("Custom!", "#FF0000",
        divoom.WithPosition(10, 20),
        divoom.WithFont(5),
        divoom.WithAlignment(1), // left align
    )
}
```

### Method 2: Manual Control (Full Control)

```go
package main

import (
    "github.com/eugene-bert/divoom-golang-api"
)

func main() {
    client := divoom.NewClient("192.168.1.180")

    // Step 1: Switch to Custom channel
    client.SetChannelIndex(3)

    // Step 2: Switch to CustomPageIndex 1 (CRITICAL!)
    client.SetCustomPageIndex(1)

    // Step 3: Send background (blank or colored)
    client.SendBlankScreen()
    // OR
    client.SendColorScreen(0x000000) // Black background

    // Step 4: Overlay text
    client.SendText(divoom.TextParams{
        TextID:     1,
        X:          0,
        Y:          24,
        Font:       4,
        TextWidth:  64,
        Speed:      0,
        TextString: "Hello!",
        Color:      "#00FF00",
        Align:      2, // center
    })
}
```

### Method 3: Custom Animation + Text

```go
package main

import (
    "encoding/base64"
    "github.com/eugene-bert/divoom-golang-api"
)

func main() {
    client := divoom.NewClient("192.168.1.180")

    // Step 1: Setup channel
    client.SetChannelIndex(3)
    client.SetCustomPageIndex(1)

    // Step 2: Create 2-frame RED/BLUE animation
    // Frame 1: RED
    redPixels := make([]byte, 64*64*3)
    for i := 0; i < 64*64; i++ {
        redPixels[i*3] = 255 // R
    }
    redData := base64.StdEncoding.EncodeToString(redPixels)

    client.SendGif(divoom.GifParams{
        PicNum:    2,
        PicWidth:  64,
        PicOffset: 0,
        PicID:     100,
        PicSpeed:  500,
        PicData:   redData,
    })

    // Frame 2: BLUE
    bluePixels := make([]byte, 64*64*3)
    for i := 0; i < 64*64; i++ {
        bluePixels[i*3+2] = 255 // B
    }
    blueData := base64.StdEncoding.EncodeToString(bluePixels)

    client.SendGif(divoom.GifParams{
        PicNum:    2,
        PicWidth:  64,
        PicOffset: 1,
        PicID:     100,
        PicSpeed:  500,
        PicData:   blueData,
    })

    // Step 3: Overlay text
    client.SendText(divoom.TextParams{
        TextID:     1,
        X:          0,
        Y:          24,
        Font:       4,
        TextWidth:  64,
        Speed:      0,
        TextString: "Animated!",
        Color:      "#FFFFFF",
        Align:      2,
    })
}
```

### Method 4: GIF from URL + Text

```go
package main

import (
    "github.com/eugene-bert/divoom-golang-api"
)

func main() {
    client := divoom.NewClient("192.168.1.180")

    // Step 1: Setup channel
    client.SetChannelIndex(3)
    client.SetCustomPageIndex(1)

    // Step 2: Play GIF from URL
    client.PlayTFGif("http://f.divoom-gz.com/64_64.gif")

    // Step 3: Overlay text
    client.SendText(divoom.TextParams{
        TextID:     1,
        X:          0,
        Y:          24,
        Font:       4,
        TextWidth:  64,
        Speed:      0,
        TextString: "URL GIF!",
        Color:      "#FFFF00",
        Align:      2,
    })
}
```

## Key Discoveries

### Why Text Wasn't Working

1. **Wrong Custom Page** - We were trying page 0 (favorites) instead of page 1 (custom content)
2. **Wrong Order** - Sending GIF before switching to the page didn't work
3. **Missing Base Layer** - Text requires a GIF/animation as a base layer (API design quirk)

### What Works

✅ Tools (timer, stopwatch, scoreboard, buzzer) - work independently
✅ Brightness, rotation, channel switching - work as expected
✅ GIFs via `PlayTFGif(url)` - display immediately
✅ GIFs via `SendHttpGif` - display when on correct page
✅ Text via `SendText` - overlays on GIFs when done correctly

### What Doesn't Work

❌ Text without a base GIF/animation
❌ `SendHttpGif` when not on CustomPageIndex 1
❌ Switching to Custom channel without specifying the page

## Testing

Run the working example:

```bash
go run cmd/working-example/main.go
```

You should see:
1. Green "Hello PIXOO64!" text
2. Red "Custom Text!" at custom position
3. Blue background with white "Blue BG!" text
4. GIF with yellow "GIF + Text!" overlay

## Architecture Notes

### Custom Channel Pages

- **Page 0 (CustomPageIndex 0)**: Displays your saved favorites from the Divoom app
- **Page 1 (CustomPageIndex 1)**: Displays custom uploaded content (use this!)
- **Page 2 (CustomPageIndex 2)**: Shows test pattern (unclear purpose)

### API Commands

- `Channel/SetIndex` - Switch channels (0-4)
- `Channel/SetCustomPageIndex` - Switch custom pages (0-2)
- `Draw/SendHttpGif` - Upload GIF to device (displays if on correct page)
- `Draw/SendHttpText` - Overlay text on current animation
- `Draw/SendHttpItemList` - Alternative text method (type: 22)
- `Device/PlayTFGif` - Play GIF from URL

## Troubleshooting

**Problem**: Text not displaying
**Solution**: Ensure you're on Channel 3, CustomPageIndex 1, and have sent a GIF first

**Problem**: GIF not displaying after SendHttpGif
**Solution**: Switch to CustomPageIndex 1 BEFORE sending the GIF

**Problem**: Shows favorites instead of custom content
**Solution**: Use CustomPageIndex 1, not 0

**Problem**: Test pattern stuck on screen
**Solution**: Send a GIF using SendHttpGif or PlayTFGif

## Success!

After this long investigation, we now have a fully working implementation for displaying text and custom animations on the PIXOO64! 🎉
