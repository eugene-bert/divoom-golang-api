package main

import (
	"fmt"
	"log"
	"time"

	divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
	// Your PIXOO64's IP address
	deviceIP := "192.168.1.180"

	// Create a new client
	client := divoom.NewClient(deviceIP)

	// Step 1: Switch to Custom channel (channel 3)
	fmt.Println("Step 1: Switching to Custom channel...")
	if err := client.SetChannelIndex(3); err != nil {
		log.Fatalf("Failed to switch channel: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// Step 2: Send blank screen (REQUIRED before text!)
	// Text must be overlaid on a GIF/animation
	fmt.Println("Step 2: Sending blank screen...")
	if err := client.SendBlankScreen(); err != nil {
		log.Fatalf("Failed to send blank screen: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// Step 3: Now send text (will overlay on the blank screen)
	fmt.Println("Step 3: Sending text...")
	err := client.SendText(divoom.TextParams{
		TextID:     1,
		X:          0,
		Y:          24,
		Direction:  0,
		Font:       4,
		TextWidth:  64,
		Speed:      0,
		TextString: "Hello PIXOO64!",
		Color:      "#00FF00",
		Align:      2,
	})
	if err != nil {
		log.Fatalf("Error sending text: %v", err)
	}

	fmt.Println()
	fmt.Println("✓✓✓ Text should now be visible on your PIXOO64! ✓✓✓")
	fmt.Println()
	fmt.Println("TIP: Use the convenience method for simpler code:")
	fmt.Println(`  client.DisplayText("Hello!", "#00FF00")`)
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  - Colors: #FF0000 (red), #0000FF (blue), #FFFF00 (yellow)")
	fmt.Println("  - Fonts: 2-7 (try different ones)")
	fmt.Println("  - Scrolling: WithScroll(100, 0) for slow scroll")
}
