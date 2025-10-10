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

	fmt.Println("New approach: Switch to empty page FIRST, THEN send GIF")
	fmt.Println()

	// Step 1: Switch to Custom channel
	fmt.Println("Step 1: Switching to Custom channel...")
	if err := client.SetChannelIndex(3); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(1 * time.Second)

	// Step 2: Go to the page that showed the test pattern (try page 1)
	fmt.Println("Step 2: Switching to CustomPageIndex 1 (test pattern page)...")
	if err := client.SetCustomPageIndex(1); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("You should see the test pattern now...")

	// Step 3: NOW send the RED GIF while we're on this page
	fmt.Println("\nStep 3: NOW sending BRIGHT RED GIF...")

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

	fmt.Println("✓ RED GIF sent!")
	fmt.Println()
	fmt.Println("════════════════════════════════════════════")
	fmt.Println("DID THE TEST PATTERN CHANGE TO RED SCREEN?")
	fmt.Println("════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("Waiting 5 seconds to observe...")
	time.Sleep(5 * time.Second)

	// Step 4: Try sending text overlay
	fmt.Println("\nStep 4: Sending WHITE text overlay...")
	if err := client.SendText(divoom.TextParams{
		TextID:     1,
		X:          0,
		Y:          24,
		Direction:  0,
		Font:       4,
		TextWidth:  64,
		Speed:      0,
		TextString: "HELLO!",
		Color:      "#FFFFFF",
		Align:      2,
	}); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("✓ Text sent!")
	fmt.Println()
	fmt.Println("DO YOU SEE 'HELLO!' IN WHITE TEXT?")
	fmt.Println()
}
