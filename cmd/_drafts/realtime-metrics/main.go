package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/eugene-bert/divoom-golang-api"
)

// simulateMetric simulates a metric value (replace with real data)
func simulateMetric(t float64) float64 {
	base := 30.0
	wave := 30.0 * math.Sin(t/10.0)
	noise := rand.Float64()*20.0 - 10.0
	value := base + wave + noise

	if value < 0 {
		value = 0
	}
	if value > 100 {
		value = 100
	}
	return value
}

// getBarChart creates a simple text bar chart
func getBarChart(value float64, maxWidth int) string {
	filled := int(value * float64(maxWidth) / 100.0)
	bar := ""
	for i := 0; i < maxWidth; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}
	return bar
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

func main() {
	// IMPORTANT: Replace with your device IP
	deviceIP := "192.168.1.180"
	client := divoom.NewClient(deviceIP)

	fmt.Println("🚀 Starting Real-Time Metrics Display on PIXOO64")
	fmt.Println("   Device IP:", deviceIP)
	fmt.Println("   Updates instantly with NO loading animations!")
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

	// Send a simple colored background ONCE
	fmt.Println("🎨 Setting up background...")
	client.SendColorScreen(0x000010) // Dark blue
	time.Sleep(500 * time.Millisecond)

	fmt.Println("📊 Starting real-time updates...")
	fmt.Println()

	// Main loop - update every 500ms (fast updates!)
	ticker := time.NewTicker(500 * time.Millisecond)
	counter := 0.0

	for range ticker.C {
		// 1. Get metric values
		cpuValue := simulateMetric(counter)
		memValue := simulateMetric(counter + 5)
		counter++

		// 2. Create display text
		cpuColor := getColor(cpuValue)
		memColor := getColor(memValue)

		// Clear all previous text
		client.ClearText()
		time.Sleep(50 * time.Millisecond)

		// Display CPU metric
		if err := client.SendText(divoom.TextParams{
			TextID:     1,
			X:          2,
			Y:          5,
			Font:       1, // Small font
			TextWidth:  60,
			TextString: "CPU",
			Color:      "#FFFFFF",
			Align:      1, // Left align
		}); err != nil {
			log.Printf("Error: %v", err)
		}

		if err := client.SendText(divoom.TextParams{
			TextID:     2,
			X:          2,
			Y:          12,
			Font:       3, // Larger font for value
			TextWidth:  60,
			TextString: fmt.Sprintf("%.1f%%", cpuValue),
			Color:      cpuColor,
			Align:      1,
		}); err != nil {
			log.Printf("Error: %v", err)
		}

		// Display bar chart for CPU
		if err := client.SendText(divoom.TextParams{
			TextID:     3,
			X:          2,
			Y:          21,
			Font:       0, // Smallest font
			TextWidth:  60,
			TextString: getBarChart(cpuValue, 10),
			Color:      cpuColor,
			Align:      1,
		}); err != nil {
			log.Printf("Error: %v", err)
		}

		// Display Memory metric
		if err := client.SendText(divoom.TextParams{
			TextID:     4,
			X:          2,
			Y:          32,
			Font:       1,
			TextWidth:  60,
			TextString: "MEM",
			Color:      "#FFFFFF",
			Align:      1,
		}); err != nil {
			log.Printf("Error: %v", err)
		}

		if err := client.SendText(divoom.TextParams{
			TextID:     5,
			X:          2,
			Y:          39,
			Font:       3,
			TextWidth:  60,
			TextString: fmt.Sprintf("%.1f%%", memValue),
			Color:      memColor,
			Align:      1,
		}); err != nil {
			log.Printf("Error: %v", err)
		}

		// Display bar chart for Memory
		if err := client.SendText(divoom.TextParams{
			TextID:     6,
			X:          2,
			Y:          48,
			Font:       0,
			TextWidth:  60,
			TextString: getBarChart(memValue, 10),
			Color:      memColor,
			Align:      1,
		}); err != nil {
			log.Printf("Error: %v", err)
		}

		// Display timestamp
		if err := client.SendText(divoom.TextParams{
			TextID:     7,
			X:          2,
			Y:          58,
			Font:       0,
			TextWidth:  60,
			TextString: time.Now().Format("15:04:05"),
			Color:      "#808080",
			Align:      1,
		}); err != nil {
			log.Printf("Error: %v", err)
		}

		// Console output
		fmt.Printf("📊 CPU: %5.1f%% [%s%s%s] | MEM: %5.1f%% [%s%s%s]\n",
			cpuValue, cpuColor, getBarChart(cpuValue, 5), "\033[0m",
			memValue, memColor, getBarChart(memValue, 5), "\033[0m")
	}
}
