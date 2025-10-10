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

	fmt.Println("Testing: SendHttpGif while ON custom page 1")
	fmt.Println()

	// Step 1: Switch to Custom channel, page 1
	fmt.Println("Step 1: Switching to Custom channel, page 1...")
	if err := client.SetChannelIndex(3); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	if err := client.SetCustomPageIndex(1); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("✓ On page 1 (test pattern)")

	// Step 2: Create and send a simple 2-frame animation
	// Frame 1: RED screen
	// Frame 2: BLUE screen
	fmt.Println("\nStep 2: Uploading 2-frame animation (RED → BLUE)...")

	// Frame 1: RED
	redPixels := make([]byte, 64*64*3)
	for i := 0; i < 64*64; i++ {
		redPixels[i*3] = 255 // R
		redPixels[i*3+1] = 0 // G
		redPixels[i*3+2] = 0 // B
	}
	redData := base64.StdEncoding.EncodeToString(redPixels)

	if err := client.SendGif(divoom.GifParams{
		PicNum:    2, // 2 frames total
		PicWidth:  64,
		PicOffset: 0,   // First frame
		PicID:     100, // New ID
		PicSpeed:  500, // 500ms per frame
		PicData:   redData,
	}); err != nil {
		log.Fatalf("Error sending frame 1: %v", err)
	}
	fmt.Println("  ✓ Frame 1 (RED) uploaded")
	time.Sleep(200 * time.Millisecond)

	// Frame 2: BLUE
	bluePixels := make([]byte, 64*64*3)
	for i := 0; i < 64*64; i++ {
		bluePixels[i*3] = 0     // R
		bluePixels[i*3+1] = 0   // G
		bluePixels[i*3+2] = 255 // B
	}
	blueData := base64.StdEncoding.EncodeToString(bluePixels)

	if err := client.SendGif(divoom.GifParams{
		PicNum:    2, // 2 frames total
		PicWidth:  64,
		PicOffset: 1,   // Second frame
		PicID:     100, // Same ID
		PicSpeed:  500,
		PicData:   blueData,
	}); err != nil {
		log.Fatalf("Error sending frame 2: %v", err)
	}
	fmt.Println("  ✓ Frame 2 (BLUE) uploaded")

	fmt.Println("\n════════════════════════════════════════════")
	fmt.Println("DO YOU SEE RED AND BLUE FLASHING?")
	fmt.Println("(Alternating every 500ms)")
	fmt.Println("════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("Waiting 5 seconds to observe...")
	time.Sleep(5 * time.Second)

	// Step 3: Try text overlay on OUR animation
	fmt.Println("\nStep 3: Sending text overlay via SendText...")
	if err := client.SendText(divoom.TextParams{
		TextID:     1,
		X:          0,
		Y:          24,
		Direction:  0,
		Font:       4,
		TextWidth:  64,
		Speed:      0,
		TextString: "SUCCESS!",
		Color:      "#FFFFFF", // White
		Align:      2,
	}); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("✓ Text sent!")
	fmt.Println()
	fmt.Println("DO YOU SEE WHITE 'SUCCESS!' TEXT")
	fmt.Println("ON THE FLASHING RED/BLUE ANIMATION?")
	fmt.Println()
}
