package main

import (
	"fmt"
	"log"
	"os"
	"time"

	divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
	ip := "192.168.1.180"
	if len(os.Args) > 1 {
		ip = os.Args[1]
	}

	client := divoom.NewClient(ip)
	fmt.Printf("Connecting to PIXOO64 at %s\n", ip)

	// Display text
	if err := client.DisplayText("Hello!", "#00FF00"); err != nil {
		log.Fatalf("DisplayText: %v", err)
	}
	fmt.Println("Text displayed")
	time.Sleep(3 * time.Second)

	// Canvas drawing
	client.ResetGifID()
	client.SetChannelIndex(3)
	client.SetCustomPageIndex(1)
	time.Sleep(2 * time.Second)

	canvas := divoom.NewCanvas(client)
	canvas.Clear(0, 0, 20)
	canvas.DrawLine(0, 0, 63, 63, 255, 0, 0)
	canvas.FillRectangle(20, 20, 43, 43, 0, 255, 0)
	canvas.DrawCircle(32, 32, 15, 255, 255, 0)

	if err := canvas.Push(); err != nil {
		log.Fatalf("Push: %v", err)
	}
	fmt.Println("Canvas pushed")
}
