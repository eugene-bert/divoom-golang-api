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

	fmt.Println("Testing Custom Page activation...")
	fmt.Println()

	// Step 1: Switch to Custom channel
	fmt.Println("Step 1: Switching to Custom channel (3)...")
	if err := client.SetChannelIndex(3); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// Step 2: Set Custom Page Index (THE MISSING PIECE!)
	fmt.Println("Step 2: Setting Custom Page Index to 0...")
	if err := client.SetCustomPageIndex(0); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// Step 3: Send a BRIGHT RED GIF
	fmt.Println("Step 3: Sending BRIGHT RED screen...")

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
	fmt.Println("✓✓✓ DO YOU SEE A BRIGHT RED SCREEN NOW? ✓✓✓")
	fmt.Println()
	fmt.Println("Waiting 5 seconds...")
	time.Sleep(5 * time.Second)

	// Step 4: Try text overlay
	fmt.Println("\nStep 4: Sending WHITE text overlay...")
	if err := client.SendText(divoom.TextParams{
		TextID:     1,
		X:          0,
		Y:          24,
		Direction:  0,
		Font:       4,
		TextWidth:  64,
		Speed:      0,
		TextString: "SUCCESS!",
		Color:      "#FFFFFF",
		Align:      2,
	}); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("✓ Text sent!")
	fmt.Println()
	fmt.Println("DO YOU SEE 'SUCCESS!' IN WHITE TEXT?")
	fmt.Println()
}
