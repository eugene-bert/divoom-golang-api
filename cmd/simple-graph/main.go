package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/eugene-bert/divoom-golang-api"
)

// simulateMetric simulates a metric value
func simulateMetric(t float64) float64 {
	base := 50.0
	wave := 30.0 * math.Sin(t/5.0)
	noise := rand.Float64()*10.0 - 5.0
	value := base + wave + noise

	if value < 0 {
		value = 0
	}
	if value > 100 {
		value = 100
	}
	return value
}

// getColor returns RGB color based on value
func getColor(value float64) (r, g, b byte) {
	if value < 50 {
		return 0, 255, 0 // Green
	} else if value < 75 {
		return 255, 255, 0 // Yellow
	}
	return 255, 0, 0 // Red
}

func main() {
	deviceIP := "192.168.1.180"
	client := divoom.NewClient(deviceIP)

	fmt.Println("🎨 Simple Graph Using Canvas API")
	fmt.Println("   Much cleaner API!")
	fmt.Println("   Updates every 2 seconds")
	fmt.Println("   Press Ctrl+C to stop")
	fmt.Println()

	// Setup device
	fmt.Println("⚙️  Setting up...")
	client.ResetGifID()
	time.Sleep(500 * time.Millisecond)
	client.SetChannelIndex(3)
	time.Sleep(500 * time.Millisecond)
	client.SetCustomPageIndex(1)
	time.Sleep(2 * time.Second)

	// Create canvas
	canvas := divoom.NewCanvas(client)

	// Keep last 60 values (almost full width)
	history := []float64{}
	maxHistory := 60

	fmt.Println("📊 Drawing graphs...\n")

	ticker := time.NewTicker(2 * time.Second)
	counter := 0.0
	picID := 1

	for range ticker.C {
		// Get new value
		value := simulateMetric(counter)
		counter++

		// Add to history
		history = append(history, value)
		if len(history) > maxHistory {
			history = history[1:]
		}

		// Clear canvas with dark background
		canvas.Clear(0, 0, 20)

		// Draw grid lines (every 12 pixels vertically)
		for y := 12; y < 64; y += 12 {
			for x := 0; x < 64; x++ {
				canvas.SetPixel(x, y, 30, 30, 40)
			}
		}

		// Draw the graph line
		for i := 0; i < len(history)-1; i++ {
			// Map data point to screen position
			x1 := i * 64 / maxHistory
			x2 := (i + 1) * 64 / maxHistory

			// Map value (0-100) to y position (10-60, inverted)
			y1 := 60 - int(history[i]*50/100)
			y2 := 60 - int(history[i+1]*50/100)

			// Get color
			r, g, b := getColor(history[i])

			// Draw line segment
			canvas.DrawLine(x1, y1, x2, y2, r, g, b)
		}

		// Draw a filled rectangle at current value position
		if len(history) > 0 {
			lastVal := history[len(history)-1]
			y := 60 - int(lastVal*50/100)
			r, g, b := getColor(lastVal)
			canvas.FillRectangle(61, y-2, 63, y+2, r, g, b)
		}

		// Push to device (will show loading animation)
		if picID > 10 {
			client.ResetGifID()
			time.Sleep(500 * time.Millisecond)
			picID = 1
		}

		// Actually we need to update canvas.Push() to use picID
		// For now, just push directly
		if err := canvas.Push(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		picID++

		// Overlay text with current value
		color := "#00FF00"
		if value >= 75 {
			color = "#FF0000"
		} else if value >= 50 {
			color = "#FFFF00"
		}

		client.SendText(divoom.TextParams{
			TextID: 1, X: 0, Y: 2, Font: 3,
			TextWidth: 64, TextString: fmt.Sprintf("%.0f%%", value),
			Color: color, Align: 2,
		})

		fmt.Printf("📊 Value: %.1f%% | Points: %d\n", value, len(history))
	}
}
