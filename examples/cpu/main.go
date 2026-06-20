package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	divoom "github.com/eugene-bert/divoom-golang-api"
	"github.com/shirou/gopsutil/v3/cpu"
)

func main() {
	ip := "192.168.2.180"
	if len(os.Args) > 1 {
		ip = os.Args[1]
	}

	client := divoom.NewClient(ip)
	fmt.Printf("CPU monitor → PIXOO64 at %s (Ctrl+C to stop)\n", ip)

	client.ResetGifID()
	client.SetChannelIndex(3)
	client.SetCustomPageIndex(1)
	time.Sleep(2 * time.Second)

	canvas := divoom.NewCanvas(client)
	history := make([]float64, 0, 60)

	// Establish CPU baseline
	cpu.Percent(0, false)

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
			pct, err := cpu.Percent(0, false)
			if err != nil || len(pct) == 0 {
				continue
			}
			val := pct[0]

			history = append(history, val)
			if len(history) > 60 {
				history = history[1:]
			}

			// Dark blue background
			canvas.Clear(0, 0, 20)

			// Grid lines
			for y := 12; y < 64; y += 12 {
				for x := 0; x < 64; x++ {
					canvas.SetPixel(x, y, 30, 30, 40)
				}
			}

			// Graph line
			for i := 0; i < len(history)-1; i++ {
				x1 := i * 64 / 60
				x2 := (i + 1) * 64 / 60
				y1 := 60 - int(history[i]*50/100)
				y2 := 60 - int(history[i+1]*50/100)
				r, g, b := cpuColor(history[i])
				canvas.DrawLine(x1, y1, x2, y2, r, g, b)
			}

			// Current value dot
			if len(history) > 0 {
				last := history[len(history)-1]
				y := 60 - int(last*50/100)
				r, g, b := cpuColor(last)
				canvas.FillRectangle(61, y-2, 63, y+2, r, g, b)
			}

			if err := canvas.Push(); err != nil {
				log.Printf("Push error: %v", err)
				continue
			}

			// Overlay percentage text
			color := "#00FF00"
			if val >= 75 {
				color = "#FF0000"
			} else if val >= 50 {
				color = "#FFFF00"
			}
			client.SendText(divoom.TextParams{
				TextID: 1, X: 2, Y: 2, Font: 3, TextWidth: 64,
				TextString: fmt.Sprintf("%3.0f%%", val),
				Color:      color, Align: 1,
			})
		}
	}
}

func cpuColor(val float64) (r, g, b byte) {
	if val < 50 {
		return 0, 255, 0
	} else if val < 75 {
		return 255, 255, 0
	}
	return 255, 0, 0
}
