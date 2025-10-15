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

// getColor returns RGB based on value
func getColor(value float64) (r, g, b byte) {
	if value < 50 {
		return 0, 255, 0
	} else if value < 75 {
		return 255, 255, 0
	}
	return 255, 0, 0
}

func main() {
	deviceIP := "192.168.1.180"
	client := divoom.NewClient(deviceIP)

	fmt.Println("🐍 Python Library Port - Exact Approach")
	fmt.Println("   - Single frame (PicNum: 1)")
	fmt.Println("   - Incrementing PicID")
	fmt.Println("   - Reset every 32 frames")
	fmt.Println("   Testing if this reduces loading...")
	fmt.Println()

	// Setup
	fmt.Println("⚙️  Setting up...")
	client.ResetGifID()
	time.Sleep(500 * time.Millisecond)
	client.SetChannelIndex(3)
	time.Sleep(500 * time.Millisecond)
	client.SetCustomPageIndex(1)
	time.Sleep(2 * time.Second)

	// Create Python-style canvas
	canvas := divoom.NewPixooCanvas(client)

	// History for graph
	history := []float64{}
	maxHistory := 60

	fmt.Println("📊 Starting updates (Python approach)...\n")

	ticker := time.NewTicker(1 * time.Second) // Try 1 second updates
	counter := 0.0
	frameCount := 0

	for range ticker.C {
		frameCount++

		// Get new value
		value := simulateMetric(counter)
		counter++

		// Add to history
		history = append(history, value)
		if len(history) > maxHistory {
			history = history[1:]
		}

		// Clear canvas
		canvas.Clear(0, 0, 20)

		// Draw grid
		for y := 12; y < 64; y += 12 {
			for x := 0; x < 64; x++ {
				canvas.SetPixel(x, y, 30, 30, 40)
			}
		}

		// Draw graph line
		for i := 0; i < len(history)-1; i++ {
			x1 := i * 64 / maxHistory
			x2 := (i + 1) * 64 / maxHistory
			y1 := 60 - int(history[i]*50/100)
			y2 := 60 - int(history[i+1]*50/100)
			r, g, b := getColor(history[i])
			canvas.DrawLine(x1, y1, x2, y2, r, g, b)
		}

		// Current value indicator
		if len(history) > 0 {
			lastVal := history[len(history)-1]
			y := 60 - int(lastVal*50/100)
			r, g, b := getColor(lastVal)
			canvas.FillRectangle(61, y-2, 63, y+2, r, g, b)
		}

		// Push using Python approach
		if err := canvas.Push(); err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		// Add text overlay
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

		fmt.Printf("📊 Frame %d | Value: %.1f%% | Will reset at frame 32\n",
			frameCount, value)
	}
}
