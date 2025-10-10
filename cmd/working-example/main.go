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

	fmt.Println("=== Divoom PIXOO64 - Working Example ===")
	fmt.Println()

	// Example 1: Display text the EASY way (recommended)
	fmt.Println("Example 1: Display text using DisplayText() helper")
	fmt.Println("This automatically handles channel switching and setup")

	if err := client.DisplayText("Hello PIXOO64!", "#00FF00"); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("✓ Green 'Hello PIXOO64!' text should be displayed")
	time.Sleep(3 * time.Second)

	// Example 2: Display text with custom options
	fmt.Println("\nExample 2: Display text with custom position and color")

	if err := client.DisplayText("Custom Text!", "#FF0000",
		divoom.WithPosition(5, 10),
		divoom.WithAlignment(1), // left align
	); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("✓ Red 'Custom Text!' at position (5,10)")
	time.Sleep(3 * time.Second)

	// Example 3: Manual approach with custom animation
	fmt.Println("\nExample 3: Custom animation with text overlay (manual)")
	fmt.Println("This shows the full control approach")

	// Step 0: Reset GIF ID (important when doing manual approach!)
	if err := client.ResetGifID(); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// Step 1: Switch to Custom channel page 1
	if err := client.SetChannelIndex(3); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	if err := client.SetCustomPageIndex(1); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(2 * time.Second)

	// Step 2: Send colored background
	if err := client.SendColorScreen(0x0000FF); err != nil { // Blue
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// Step 3: Overlay text
	if err := client.SendText(divoom.TextParams{
		TextID:     1,
		X:          0,
		Y:          24,
		Font:       4,
		TextWidth:  64,
		Speed:      0,
		TextString: "Blue BG!",
		Color:      "#FFFFFF",
		Align:      2,
	}); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("✓ Blue background with white 'Blue BG!' text")
	time.Sleep(3 * time.Second)

	// Example 4: Play GIF from URL with text overlay
	fmt.Println("\nExample 4: Play GIF from URL with text overlay")

	// Note: PlayTFGif doesn't need ResetGifID since it plays from URL
	if err := client.SetChannelIndex(3); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	if err := client.SetCustomPageIndex(1); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(2 * time.Second)

	// Play a GIF from Divoom's server
	if err := client.PlayTFGif("http://f.divoom-gz.com/64_64.gif"); err != nil {
		log.Fatalf("Error: %v", err)
	}

	time.Sleep(1 * time.Second)

	// Overlay text on the GIF
	if err := client.SendText(divoom.TextParams{
		TextID:     1,
		X:          0,
		Y:          24,
		Font:       4,
		TextWidth:  64,
		Speed:      0,
		TextString: "GIF + Text!",
		Color:      "#FFFF00",
		Align:      2,
	}); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("✓ GIF playing with yellow 'GIF + Text!' overlay")

	fmt.Println("\n=== All examples complete! ===")
}
