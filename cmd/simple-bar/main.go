package main

import (
	"fmt"
	"log"
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

// getColor returns color based on value
func getColor(value float64) string {
	if value < 50 {
		return "#00FF00" // Green
	} else if value < 75 {
		return "#FFFF00" // Yellow
	}
	return "#FF0000" // Red
}

// createBar creates a simple horizontal bar
func createBar(value float64, maxChars int) string {
	filled := int(value * float64(maxChars) / 100.0)
	bar := ""
	for i := 0; i < maxChars; i++ {
		if i < filled {
			bar += "#"
		} else {
			bar += "."
		}
	}
	return bar
}

func main() {
	// IMPORTANT: Replace with your device IP
	deviceIP := "192.168.1.180"
	client := divoom.NewClient(deviceIP)

	fmt.Println("🚀 Simple Real-Time Bar Graph on PIXOO64")
	fmt.Println("   Device IP:", deviceIP)
	fmt.Println("   Press Ctrl+C to stop")
	fmt.Println()

	// Setup device ONCE
	fmt.Println("⚙️  Setting up device...")
	client.ResetGifID()
	time.Sleep(500 * time.Millisecond)
	client.SetChannelIndex(3)
	time.Sleep(500 * time.Millisecond)
	client.SetCustomPageIndex(1)
	time.Sleep(2 * time.Second)

	// Send dark background ONCE
	fmt.Println("🎨 Setting up background...")
	client.SendColorScreen(0x000000) // Black
	time.Sleep(500 * time.Millisecond)

	fmt.Println("📊 Starting updates...")
	fmt.Println()

	// Keep last 3 values for mini history
	history := []float64{}
	maxHistory := 3

	// Main loop - update every 1 second
	ticker := time.NewTicker(1 * time.Second)
	counter := 0.0

	for range ticker.C {
		// Get new value
		value := simulateMetric(counter)
		counter++

		// Add to history
		history = append(history, value)
		if len(history) > maxHistory {
			history = history[1:]
		}

		// Clear previous text
		client.ClearText()
		time.Sleep(50 * time.Millisecond)

		// Title
		if err := client.SendText(divoom.TextParams{
			TextID:     1,
			X:          0,
			Y:          2,
			Font:       2,
			TextWidth:  64,
			TextString: "CPU LOAD",
			Color:      "#FFFFFF",
			Align:      2, // Center
		}); err != nil {
			log.Printf("Error: %v", err)
		}

		// Current value - BIG
		valueColor := getColor(value)
		if err := client.SendText(divoom.TextParams{
			TextID:     2,
			X:          0,
			Y:          15,
			Font:       5, // Largest font!
			TextWidth:  64,
			TextString: fmt.Sprintf("%.0f%%", value),
			Color:      valueColor,
			Align:      2, // Center
		}); err != nil {
			log.Printf("Error: %v", err)
		}

		// Bar chart - current value
		bar := createBar(value, 10)
		if err := client.SendText(divoom.TextParams{
			TextID:     3,
			X:          0,
			Y:          32,
			Font:       3, // Medium font
			TextWidth:  64,
			TextString: bar,
			Color:      valueColor,
			Align:      2, // Center
		}); err != nil {
			log.Printf("Error: %v", err)
		}

		// Show last 3 values as mini bars
		yPos := 42
		for i, histVal := range history {
			histColor := getColor(histVal)
			histBar := createBar(histVal, 8)

			if err := client.SendText(divoom.TextParams{
				TextID:     4 + i,
				X:          0,
				Y:          yPos,
				Font:       1, // Small font
				TextWidth:  64,
				TextString: histBar,
				Color:      histColor,
				Align:      2, // Center
			}); err != nil {
				log.Printf("Error: %v", err)
			}
			yPos += 6
		}

		// Console output
		fmt.Printf("📊 %5.1f%% | Bar: [%s]\n", value, bar)
	}
}
