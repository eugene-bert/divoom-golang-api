package main

import (
	"fmt"
	"log"
	"time"

	divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
	deviceIP := "192.168.1.180"
	client := divoom.NewClient(deviceIP)

	fmt.Println("Testing: PlayTFGif + ItemList text overlay")
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
	time.Sleep(1 * time.Second)

	// Step 2: Play the Iron Man/Hulk GIF
	fmt.Println("Step 2: Playing GIF from URL...")
	if err := client.PlayTFGif("http://f.divoom-gz.com/64_64.gif"); err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("✓ GIF playing (Iron Man/Hulk)")
	time.Sleep(2 * time.Second)

	// Step 3: Try ItemList text overlay
	fmt.Println("\nStep 3: Sending text via ItemList API...")
	items := []divoom.ItemListParams{
		{
			TextID:     1,
			Type:       22, // TEXT_MESSAGE
			X:          0,
			Y:          24,
			Direction:  0,
			Font:       4,
			TextWidth:  64,
			TextHeight: 16,
			Speed:      0,
			Align:      2,
			TextString: "HELLO PIXOO!",
			Color:      "#FFFF00", // Yellow
		},
	}

	if err := client.SendHttpItemList(items); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("✓ ItemList text sent!")
	fmt.Println()
	fmt.Println("════════════════════════════════════════════")
	fmt.Println("DO YOU SEE YELLOW 'HELLO PIXOO!' TEXT")
	fmt.Println("OVERLAID ON THE IRON MAN/HULK GIF?")
	fmt.Println("════════════════════════════════════════════")
	fmt.Println()
}
