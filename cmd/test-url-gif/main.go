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

	fmt.Println("Testing PlayTFGif with URL...")
	fmt.Println()

	// Step 1: Switch to Custom channel
	fmt.Println("Step 1: Switching to Custom channel...")
	if err := client.SetChannelIndex(3); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// Step 2: Set Custom Page Index
	fmt.Println("Step 2: Setting Custom Page Index to 0...")
	if err := client.SetCustomPageIndex(0); err != nil {
		log.Fatalf("Error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// Step 3: Play a GIF from Divoom's server
	fmt.Println("Step 3: Playing GIF from URL (Divoom's test GIF)...")
	if err := client.PlayTFGif("http://f.divoom-gz.com/64_64.gif"); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("✓ PlayTFGif command sent!")
	fmt.Println()
	fmt.Println("DO YOU SEE A GIF PLAYING FROM THE URL?")
	fmt.Println("(This is a test GIF from Divoom's server)")
	fmt.Println()
}
