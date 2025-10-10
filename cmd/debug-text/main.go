package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
	deviceIP := "192.168.1.180"
	client := divoom.NewClient(deviceIP)

	// Step 1: Check current channel
	fmt.Println("Step 1: Getting current channel...")
	currentChannel, err := client.GetChannelIndex()
	if err != nil {
		log.Printf("Error getting channel: %v\n", err)
	} else {
		fmt.Printf("  Current channel: %d\n", currentChannel)
	}

	// Step 2: Switch to Custom channel
	fmt.Println("\nStep 2: Switching to Custom channel (3)...")
	if err := client.SetChannelIndex(3); err != nil {
		log.Fatalf("  ERROR: %v\n", err)
	}
	fmt.Println("  ✓ Channel switch command sent")
	time.Sleep(1 * time.Second)

	// Step 3: Verify channel switched
	fmt.Println("\nStep 3: Verifying channel switch...")
	currentChannel, err = client.GetChannelIndex()
	if err != nil {
		log.Printf("  Error: %v\n", err)
	} else {
		fmt.Printf("  Current channel: %d", currentChannel)
		if currentChannel == 3 {
			fmt.Println(" ✓ (Custom)")
		} else {
			fmt.Println(" ✗ (NOT Custom!)")
		}
	}

	// Step 4: Get all config
	fmt.Println("\nStep 4: Getting device config...")
	config, err := client.GetAllConf()
	if err != nil {
		log.Printf("  Error: %v\n", err)
	} else {
		configJSON, _ := json.MarshalIndent(config, "  ", "  ")
		fmt.Printf("  Config: %s\n", configJSON)
	}

	// Step 5: Send blank screen
	fmt.Println("\nStep 5: Sending blank screen...")
	if err := client.SendBlankScreen(); err != nil {
		log.Fatalf("  ERROR: %v\n", err)
	}
	fmt.Println("  ✓ Blank screen sent")
	time.Sleep(1 * time.Second)

	// Step 6: Send very simple text
	fmt.Println("\nStep 6: Sending text...")
	err = client.SendText(divoom.TextParams{
		TextID:     1,
		X:          0,
		Y:          0,
		Direction:  0,
		Font:       2, // Smallest font
		TextWidth:  64,
		Speed:      0,
		TextString: "TEST",
		Color:      "#FFFFFF", // White
		Align:      1,         // Left align
	})
	if err != nil {
		log.Fatalf("  ERROR: %v\n", err)
	}
	fmt.Println("  ✓ Text command sent")

	// Step 7: Try different Y positions
	fmt.Println("\nStep 7: Trying different Y positions...")
	for y := 0; y <= 48; y += 16 {
		fmt.Printf("  Sending text at Y=%d...\n", y)
		client.SendText(divoom.TextParams{
			TextID:     1,
			X:          0,
			Y:          y,
			Direction:  0,
			Font:       4,
			TextWidth:  64,
			Speed:      0,
			TextString: fmt.Sprintf("Y=%d", y),
			Color:      "#FF0000",
			Align:      1,
		})
		time.Sleep(2 * time.Second)
	}

	fmt.Println("\n✓ Debug sequence complete!")
	fmt.Println("\nDid you see ANY text appear on the display?")
	fmt.Println("If not, the issue might be:")
	fmt.Println("  1. Text requires animation to be 'playing' (not just uploaded)")
	fmt.Println("  2. Different firmware version with different behavior")
	fmt.Println("  3. Missing some initialization command")
}
