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

	fmt.Println("Testing if GIF actually displays...")
	fmt.Println()

	// Step 1: Switch to Custom channel
	fmt.Println("Step 1: Switching to Custom channel...")
	if err := client.SetChannelIndex(3); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(1 * time.Second)

	// Step 2: Send a VISIBLE colored GIF (red screen)
	fmt.Println("Step 2: Sending RED screen GIF...")

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

	fmt.Println("✓ RED screen GIF sent")
	fmt.Println()
	fmt.Println("DO YOU SEE A RED SCREEN?")
	fmt.Println("If YES: GIFs work, problem is with text overlay")
	fmt.Println("If NO: GIFs aren't displaying at all")
	fmt.Println()
	fmt.Println("Waiting 5 seconds...")
	time.Sleep(5 * time.Second)

	// Step 3: Try sending text on the red screen
	fmt.Println("\nStep 3: Sending WHITE text on red screen...")
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

	fmt.Println("✓ Text sent")
	fmt.Println()
	fmt.Println("DO YOU SEE WHITE TEXT ON RED BACKGROUND?")
	fmt.Println()
}
