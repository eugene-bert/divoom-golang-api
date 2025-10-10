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

	fmt.Println("Testing: Reset GIF ID, then try RED/BLUE animation")
	fmt.Println()

	// Step 0: Reset GIF ID counter (clear all uploaded GIFs)
	fmt.Println("Step 0: Resetting GIF ID counter...")
	if err := client.ResetGifID(); err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("✓ GIF ID reset")
	time.Sleep(1 * time.Second)

	// Step 1: Switch to Custom channel, page 1
	fmt.Println("\nStep 1: Switching to Custom channel, page 1...")
	if err := client.SetChannelIndex(3); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	if err := client.SetCustomPageIndex(1); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("✓ On page 1 (test pattern)")

	// Step 2: Create RED/BLUE animation with PicID 1 (since we just reset)
	fmt.Println("\nStep 2: Uploading 2-frame RED/BLUE animation...")

	// Frame 1: RED
	redPixels := make([]byte, 64*64*3)
	for i := 0; i < 64*64; i++ {
		redPixels[i*3] = 255 // R
	}
	redData := base64.StdEncoding.EncodeToString(redPixels)

	if err := client.SendGif(divoom.GifParams{
		PicNum:    2,
		PicWidth:  64,
		PicOffset: 0,
		PicID:     1, // Start from 1 since we reset
		PicSpeed:  500,
		PicData:   redData,
	}); err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("  ✓ Frame 1 (RED) uploaded")
	time.Sleep(200 * time.Millisecond)

	// Frame 2: BLUE
	bluePixels := make([]byte, 64*64*3)
	for i := 0; i < 64*64; i++ {
		bluePixels[i*3+2] = 255 // B
	}
	blueData := base64.StdEncoding.EncodeToString(bluePixels)

	if err := client.SendGif(divoom.GifParams{
		PicNum:    2,
		PicWidth:  64,
		PicOffset: 1,
		PicID:     1,
		PicSpeed:  500,
		PicData:   blueData,
	}); err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("  ✓ Frame 2 (BLUE) uploaded")

	fmt.Println()
	fmt.Println("════════════════════════════════════════════")
	fmt.Println("DO YOU SEE RED AND BLUE FLASHING?")
	fmt.Println("════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("Waiting 5 seconds...")
	time.Sleep(5 * time.Second)

	// Step 3: Try text overlay
	fmt.Println("\nStep 3: Sending WHITE 'SUCCESS!' text overlay...")
	if err := client.SendText(divoom.TextParams{
		TextID:     1,
		X:          0,
		Y:          24,
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
	fmt.Println("DO YOU SEE WHITE 'SUCCESS!' TEXT?")
	fmt.Println()
}
