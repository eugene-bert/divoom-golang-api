package main

import (
	"encoding/base64"
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

// getColor returns color based on value
func getColor(value float64) string {
	if value < 50 {
		return "#00FF00" // Green
	} else if value < 75 {
		return "#FFFF00" // Yellow
	}
	return "#FF0000" // Red
}

// createGridBackground creates a simple grid background ONCE
func createGridBackground() []byte {
	width := 64
	height := 64
	pixels := make([]byte, width*height*3)

	bgColor := [3]byte{5, 5, 15}      // Very dark blue
	gridColor := [3]byte{30, 30, 40} // Slightly lighter for grid lines

	// Fill background
	for i := 0; i < width*height; i++ {
		pixels[i*3] = bgColor[0]
		pixels[i*3+1] = bgColor[1]
		pixels[i*3+2] = bgColor[2]
	}

	// Draw horizontal grid lines (every 10 pixels)
	for y := 10; y < height; y += 10 {
		for x := 0; x < width; x++ {
			idx := (y*width + x) * 3
			pixels[idx] = gridColor[0]
			pixels[idx+1] = gridColor[1]
			pixels[idx+2] = gridColor[2]
		}
	}

	// Draw vertical lines for bar positions
	for barX := 8; barX < width; barX += 8 {
		for y := 10; y < 60; y++ {
			idx := (y*width + barX) * 3
			pixels[idx] = gridColor[0]
			pixels[idx+1] = gridColor[1]
			pixels[idx+2] = gridColor[2]
		}
	}

	return pixels
}

// createTextBar creates vertical bar using stacked text
func createTextBar(value float64, barNum int) string {
	// Calculate how many "levels" this bar should be (0-5)
	levels := int(value / 20.0) // 0-100 maps to 0-5 levels

	// Create bar character based on level
	if levels == 0 {
		return "."
	}
	return "#"
}

func main() {
	// IMPORTANT: Replace with your device IP
	deviceIP := "192.168.1.180"
	client := divoom.NewClient(deviceIP)

	fmt.Println("🚀 Hybrid Bar Chart on PIXOO64")
	fmt.Println("   Grid background (GIF once) + Text bars (instant updates)")
	fmt.Println("   Press Ctrl+C to stop")
	fmt.Println()

	// Setup device
	fmt.Println("⚙️  Setting up device...")
	client.ResetGifID()
	time.Sleep(500 * time.Millisecond)
	client.SetChannelIndex(3)
	time.Sleep(500 * time.Millisecond)
	client.SetCustomPageIndex(1)
	time.Sleep(2 * time.Second)

	// Send grid background ONCE
	fmt.Println("🎨 Drawing grid background (one time only)...")
	pixels := createGridBackground()
	pixels2 := make([]byte, len(pixels))
	copy(pixels2, pixels)
	for i := 0; i < len(pixels2); i++ {
		if pixels2[i] < 254 {
			pixels2[i]++
		}
	}

	frame1Data := base64.StdEncoding.EncodeToString(pixels)
	frame2Data := base64.StdEncoding.EncodeToString(pixels2)

	client.SendGif(divoom.GifParams{
		PicNum: 2, PicWidth: 64, PicOffset: 0,
		PicID: 1, PicSpeed: 500, PicData: frame1Data,
	})
	time.Sleep(100 * time.Millisecond)

	client.SendGif(divoom.GifParams{
		PicNum: 2, PicWidth: 64, PicOffset: 1,
		PicID: 1, PicSpeed: 500, PicData: frame2Data,
	})
	time.Sleep(500 * time.Millisecond)

	fmt.Println("📊 Starting instant bar updates...")
	fmt.Println()

	// Keep last 7 values
	history := []float64{}
	maxBars := 7

	// Main loop - update every 500ms (FAST!)
	ticker := time.NewTicker(500 * time.Millisecond)
	counter := 0.0

	for range ticker.C {
		// Get new value
		value := simulateMetric(counter)
		counter++

		// Add to history
		history = append(history, value)
		if len(history) > maxBars {
			history = history[1:]
		}

		// Clear previous text
		client.ClearText()
		time.Sleep(30 * time.Millisecond)

		textID := 1

		// Display current value at top
		valueColor := getColor(value)
		client.SendText(divoom.TextParams{
			TextID: textID, X: 0, Y: 2, Font: 3,
			TextWidth: 64, TextString: fmt.Sprintf("%.0f%%", value),
			Color: valueColor, Align: 2,
		})
		textID++

		// Draw bars using text - 5 height levels
		for level := 0; level < 5; level++ {
			// Y position from bottom to top
			yPos := 50 - (level * 8)

			barLine := ""
			for i, histVal := range history {
				barHeight := int(histVal / 20.0) // 0-5

				if barHeight > level {
					barLine += "#"
				} else {
					barLine += " "
				}

				// Add space between bars
				if i < len(history)-1 {
					barLine += " "
				}
			}

			// Determine color for this level (based on height threshold)
			levelColor := "#00FF00"
			if level >= 3 {
				levelColor = "#FFFF00"
			}
			if level >= 4 {
				levelColor = "#FF0000"
			}

			client.SendText(divoom.TextParams{
				TextID: textID, X: 8, Y: yPos, Font: 3,
				TextWidth: 56, TextString: barLine,
				Color: levelColor, Align: 1,
			})
			textID++
		}

		// Console output
		fmt.Printf("📊 %.1f%% | Bars: ", value)
		for _, v := range history {
			fmt.Printf("%.0f ", v)
		}
		fmt.Println()
	}
}
