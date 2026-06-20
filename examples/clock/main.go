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

	// Setup: send background once, then only update text
	if err := client.DisplayText("", "#000000"); err != nil {
		fmt.Fprintf(os.Stderr, "Setup error: %v\n", err)
		os.Exit(1)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	lastMin := ""

	for {
		select {
		case <-sig:
			fmt.Println("\nStopped.")
			return
		case <-ticker.C:
			now := time.Now()
			curMin := now.Format("15:04")

			// Only update time+date when minute changes
			if curMin != lastMin {
				client.SendText(divoom.TextParams{
					TextID: 1, X: 0, Y: 10, Font: 4, TextWidth: 64,
					TextString: curMin,
					Color: "#00CCFF", Align: 2,
				})

				client.SendText(divoom.TextParams{
					TextID: 3, X: 0, Y: 44, Font: 2, TextWidth: 64,
					TextString: now.Format("02 Jan 2006"),
					Color: "#445566", Align: 2,
				})
				lastMin = curMin
			}

			// Seconds update every tick — just one SendText call
			client.SendText(divoom.TextParams{
				TextID: 2, X: 0, Y: 28, Font: 2, TextWidth: 64,
				TextString: fmt.Sprintf(":%02d", now.Second()),
				Color: "#0088AA", Align: 2,
			})
		}
	}
}
