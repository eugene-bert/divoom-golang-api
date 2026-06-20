package main

import (
	"fmt"
	"os"

	divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <device-ip> <text> [color]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s 192.168.1.100 \"Hello World!\"\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s 192.168.1.100 \"Breaking News\" FF0000\n", os.Args[0])
		os.Exit(1)
	}

	ip := os.Args[1]
	text := os.Args[2]
	color := "#00FF00"
	if len(os.Args) > 3 {
		color = divoom.ParseHexColor(os.Args[3])
	}

	client := divoom.NewClient(ip)
	fmt.Printf("Scrolling \"%s\" on PIXOO64 at %s\n", text, ip)

	err := client.DisplayText(text, color,
		divoom.WithPosition(0, 24),
		divoom.WithFont(4),
		divoom.WithScroll(50, 0),
		divoom.WithAlignment(2),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Text displayed!")
}
