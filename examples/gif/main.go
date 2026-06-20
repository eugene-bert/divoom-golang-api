package main

import (
	"fmt"
	"log"
	"os"

	divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <device-ip> <gif-path>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nExample:\n")
		fmt.Fprintf(os.Stderr, "  %s 192.168.1.100 animation.gif\n", os.Args[0])
		os.Exit(1)
	}

	ip := os.Args[1]
	gifPath := os.Args[2]

	client := divoom.NewClient(ip)
	fmt.Printf("Connecting to PIXOO64 at %s\n", ip)

	if err := client.SetChannelIndex(3); err != nil {
		log.Fatalf("Failed to set channel: %v", err)
	}
	if err := client.SetCustomPageIndex(1); err != nil {
		log.Fatalf("Failed to set page: %v", err)
	}

	fmt.Printf("Playing GIF: %s\n", gifPath)
	if err := client.PlayLocalGif(gifPath); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("GIF sent!")
}
