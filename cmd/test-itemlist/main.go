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

	fmt.Println("Testing Draw/SendHttpItemList command...")
	fmt.Println()

	// Switch to Custom channel
	fmt.Println("Step 1: Switching to Custom channel...")
	if err := client.SetChannelIndex(3); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// Send blank screen
	fmt.Println("Step 2: Sending blank screen...")
	if err := client.SendBlankScreen(); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// Method 1: Using SendHttpItemList directly
	fmt.Println("Step 3: Sending text via ItemList...")
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
			TextString: "ItemList Test!",
			Color:      "#FFFF00", // Yellow
		},
	}

	if err := client.SendHttpItemList(items); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println()
	fmt.Println("✓✓✓ ItemList command sent! ✓✓✓")
	fmt.Println()
	fmt.Println("Do you see yellow text 'ItemList Test!' on the display?")
	fmt.Println()
	fmt.Println("Waiting 5 seconds...")
	time.Sleep(5 * time.Second)

	// Try the convenience method
	fmt.Println("\nStep 4: Trying convenience method...")
	if err := client.DisplayTextViaItemList("Hello PIXOO64!", "#00FF00"); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("✓ Done!")
	fmt.Println("\nDid EITHER test display text?")
}
