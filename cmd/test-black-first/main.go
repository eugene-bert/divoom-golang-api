package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"time"

	divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
	deviceIP := "192.168.1.180"
	client := divoom.NewClient(deviceIP)

	fmt.Println("Testing: Black screen → Custom channel → Content")
	fmt.Println()

	// Step 1: Switch to BLACK SCREEN channel to stop everything
	fmt.Println("Step 1: Switching to BLACK SCREEN channel (4) to stop all content...")
	if err := client.SetChannelIndex(4); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("✓ Should see BLACK screen now")
	fmt.Println()

	// Step 2: Switch to Custom channel
	fmt.Println("Step 2: Switching to Custom channel (3)...")
	if err := client.SetChannelIndex(3); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// Step 3: Set Custom Page Index
	fmt.Println("Step 3: Setting Custom Page Index to 0...")
	if err := client.SetCustomPageIndex(0); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// Step 4: Send a BRIGHT RED screen
	fmt.Println("Step 4: Sending BRIGHT RED screen...")

	// Create 64x64 red pixels (RGB: 255, 0, 0)
	redPixels := make([]byte, 64*64*3)
	for i := 0; i < 64*64; i++ {
		redPixels[i*3] = 255 // R
		redPixels[i*3+1] = 0 // G
		redPixels[i*3+2] = 0 // B
	}
	base64Data := base64.StdEncoding.EncodeToString(redPixels)

	if err := client.SendGif(divoom.GifParams{
		PicNum:    1,
		PicWidth:  64,
		PicOffset: 0,
		PicID:     1,
		PicSpeed:  1000,
		PicData:   base64Data,
	}); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("✓ RED screen sent!")
	fmt.Println()
	fmt.Println("════════════════════════════════════════════")
	fmt.Println("DO YOU SEE A BRIGHT RED SCREEN?")
	fmt.Println("════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("Waiting 5 seconds...")
	time.Sleep(5 * time.Second)

	// Step 5: Try text overlay
	fmt.Println("\nStep 5: Sending WHITE text 'HELLO' overlay...")
	if err := client.SendText(divoom.TextParams{
		TextID:     1,
		X:          0,
		Y:          24,
		Direction:  0,
		Font:       4,
		TextWidth:  64,
		Speed:      0,
		TextString: "HELLO",
		Color:      "#FFFFFF",
		Align:      2,
	}); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("✓ Text sent!")
	fmt.Println()
	fmt.Println("DO YOU SEE WHITE 'HELLO' TEXT ON RED BACKGROUND?")
	fmt.Println()
}
