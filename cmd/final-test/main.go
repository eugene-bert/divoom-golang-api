package main

import (
	"fmt"
	"log"

	divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
	deviceIP := "192.168.1.180"
	client := divoom.NewClient(deviceIP)

	fmt.Println("Final Test: One-line text display")
	fmt.Println()

	// THE SOLUTION - ONE LINE!
	if err := client.DisplayText("SUCCESS!", "#00FF00"); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("✓✓✓ CHECK YOUR PIXOO64! ✓✓✓")
	fmt.Println()
	fmt.Println("You should see green 'SUCCESS!' text on a black background")
	fmt.Println()
	fmt.Println("If you see it - WE DID IT! 🎉")
}
