package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"time"

	divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
	deviceIP := "192.168.1.180"
	client := divoom.NewClient(deviceIP)

	fmt.Println("Testing ALL Custom Page Indexes (0, 1, 2)")
	fmt.Println()

	// Step 1: Send a BRIGHT RED GIF first
	fmt.Println("Step 1: Uploading BRIGHT RED GIF to device...")

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
	fmt.Println("✓ RED GIF uploaded")
	time.Sleep(1 * time.Second)

	// Step 2: Switch to Custom channel
	fmt.Println("\nStep 2: Switching to Custom channel...")
	if err := client.SetChannelIndex(3); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(1 * time.Second)

	// Test CustomPageIndex 0
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Testing CustomPageIndex = 0")
	fmt.Println(strings.Repeat("=", 50))
	if err := client.SetCustomPageIndex(0); err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("LOOK AT YOUR PIXOO64 NOW!")
	fmt.Println("What do you see?")
	time.Sleep(5 * time.Second)

	// Test CustomPageIndex 1
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Testing CustomPageIndex = 1")
	fmt.Println(strings.Repeat("=", 50))
	if err := client.SetCustomPageIndex(1); err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("LOOK AT YOUR PIXOO64 NOW!")
	fmt.Println("What do you see?")
	time.Sleep(5 * time.Second)

	// Test CustomPageIndex 2
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Testing CustomPageIndex = 2")
	fmt.Println(strings.Repeat("=", 50))
	if err := client.SetCustomPageIndex(2); err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("LOOK AT YOUR PIXOO64 NOW!")
	fmt.Println("What do you see?")
	time.Sleep(5 * time.Second)

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("Test complete!")
	fmt.Println("Please tell me what you saw for each CustomPageIndex (0, 1, 2)")
}
