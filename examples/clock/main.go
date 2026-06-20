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

	client.ResetGifID()
	client.SetChannelIndex(3)
	client.SetCustomPageIndex(1)
	time.Sleep(2 * time.Second)

	canvas := divoom.NewCanvas(client)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-sig:
			fmt.Println("\nStopped.")
			return
		case <-ticker.C:
			now := time.Now()

			// Dark background
			canvas.Clear(5, 5, 15)

			// Border
			canvas.DrawLine(0, 0, 63, 0, 0, 80, 160)
			canvas.DrawLine(0, 63, 63, 63, 0, 80, 160)
			canvas.DrawLine(0, 0, 0, 63, 0, 80, 160)
			canvas.DrawLine(63, 0, 63, 63, 0, 80, 160)

			// Decorative dots
			for x := 4; x < 64; x += 8 {
				canvas.SetPixel(x, 0, 0, 150, 255)
				canvas.SetPixel(x, 63, 0, 150, 255)
			}

			canvas.Push()

			// Time (large)
			client.SendText(divoom.TextParams{
				TextID: 1, X: 0, Y: 14, Font: 4, TextWidth: 64,
				TextString: now.Format("15:04"),
				Color: "#00CCFF", Align: 2,
			})

			// Seconds
			client.SendText(divoom.TextParams{
				TextID: 2, X: 0, Y: 32, Font: 2, TextWidth: 64,
				TextString: fmt.Sprintf(":%02d", now.Second()),
				Color: "#0088AA", Align: 2,
			})

			// Date
			client.SendText(divoom.TextParams{
				TextID: 3, X: 0, Y: 46, Font: 2, TextWidth: 64,
				TextString: now.Format("02 Jan 2006"),
				Color: "#445566", Align: 2,
			})
		}
	}
}
