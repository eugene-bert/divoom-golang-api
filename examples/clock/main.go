package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
	ip := "192.168.2.180"
	if len(os.Args) > 1 {
		ip = os.Args[1]
	}

	client := divoom.NewClient(ip)
	fmt.Printf("Clock → PIXOO64 at %s (Ctrl+C to stop)\n", ip)

	if err := client.DisplayText("", "#000000"); err != nil {
		fmt.Fprintf(os.Stderr, "Setup error: %v\n", err)
		os.Exit(1)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-sig:
			client.ClearText()
			fmt.Println("\nStopped.")
			return
		case <-ticker.C:
			now := time.Now()

			// Single text: HH:MM:SS
			client.SendText(divoom.TextParams{
				TextID: 1, X: 0, Y: 20, Font: 4, TextWidth: 64,
				TextString: now.Format("15:04:05"),
				Color:      "#00CCFF", Align: 2,
			})
		}
	}
}
