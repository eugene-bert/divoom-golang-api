package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <device-ip> <image-path-or-url>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s 192.168.1.100 photo.jpg\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s 192.168.1.100 https://placekitten.com/200/200\n", os.Args[0])
		os.Exit(1)
	}

	ip := os.Args[1]
	source := os.Args[2]

	client := divoom.NewClient(ip)
	fmt.Printf("Connecting to PIXOO64 at %s\n", ip)

	var err error
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		fmt.Printf("Fetching image from %s\n", source)
		err = client.DisplayImageURL(source)
	} else {
		fmt.Printf("Loading image from %s\n", source)
		err = client.DisplayImageFile(source)
	}

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Image displayed!")
}
