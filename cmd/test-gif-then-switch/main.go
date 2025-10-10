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

	fmt.Println("Testing: Send GIF FIRST, then switch to custom channel")
	fmt.Println()

	// Step 1: Black screen
	fmt.Println("Step 1: Switching to BLACK SCREEN channel (4)...")
	if err := client.SetChannelIndex(4); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("✓ Black screen active")

	// Step 2: Send BRIGHT RED GIF while still on black screen
	fmt.Println("\nStep 2: Sending BRIGHT RED GIF (while on black screen)...")

	// Create 64x64 red pixels
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
	fmt.Println("✓ RED GIF sent (to device memory)")
	time.Sleep(1 * time.Second)

	// Step 3: NOW switch to Custom channel
	fmt.Println("\nStep 3: NOW switching to Custom channel (3)...")
	if err := client.SetChannelIndex(3); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// Step 4: Set Custom Page Index
	fmt.Println("Step 4: Setting Custom Page Index to 0...")
	if err := client.SetCustomPageIndex(0); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println()
	fmt.Println("════════════════════════════════════════════")
	fmt.Println("DO YOU SEE A BRIGHT RED SCREEN NOW?")
	fmt.Println("════════════════════════════════════════════")
	fmt.Println()
}
