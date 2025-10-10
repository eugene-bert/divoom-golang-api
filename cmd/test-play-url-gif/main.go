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

	fmt.Println("Testing PlayTFGif - plays GIF from URL")
	fmt.Println()

	// Step 1: Switch to Custom channel, page 1 (test pattern page)
	fmt.Println("Step 1: Switching to Custom channel, page 1...")
	if err := client.SetChannelIndex(3); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	if err := client.SetCustomPageIndex(1); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("✓ Should see test pattern")

	// Step 2: Play a GIF from Divoom's server
	fmt.Println("\nStep 2: Playing GIF from Divoom's test URL...")
	fmt.Println("URL: http://f.divoom-gz.com/64_64.gif")

	if err := client.PlayTFGif("http://f.divoom-gz.com/64_64.gif"); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("✓ PlayTFGif command sent!")
	fmt.Println()
	fmt.Println("════════════════════════════════════════════")
	fmt.Println("DID THE TEST PATTERN CHANGE TO A GIF?")
	fmt.Println("(Should be an animation from Divoom's server)")
	fmt.Println("════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("Waiting 5 seconds...")
	time.Sleep(5 * time.Second)

	// Step 3: Try text overlay on the playing GIF
	fmt.Println("\nStep 3: Trying to overlay text on the GIF...")
	if err := client.SendText(divoom.TextParams{
		TextID:     1,
		X:          0,
		Y:          24,
		Direction:  0,
		Font:       4,
		TextWidth:  64,
		Speed:      0,
		TextString: "OVERLAY!",
		Color:      "#FFFF00", // Yellow
		Align:      2,
	}); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("✓ Text sent!")
	fmt.Println()
	fmt.Println("DO YOU SEE YELLOW 'OVERLAY!' TEXT?")
	fmt.Println()
}
