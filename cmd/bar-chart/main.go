package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/eugene-bert/divoom-golang-api"
)

const (
	width     = 64
	height    = 64
	numBars   = 8  // Number of bars to show
	barWidth  = 6  // Width of each bar in pixels
	barSpacing = 2 // Space between bars
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

// getBarColor returns RGB color based on value
func getBarColor(value float64) [3]byte {
	if value < 50 {
		return [3]byte{0, 255, 0} // Green
	} else if value < 75 {
		return [3]byte{255, 255, 0} // Yellow
	}
	return [3]byte{255, 0, 0} // Red
}

// drawBarChart draws vertical bars
func drawBarChart(values []float64) []byte {
	pixels := make([]byte, width*height*3)

	// Fill with dark background
	bgColor := [3]byte{0, 0, 20} // Very dark blue
	for i := 0; i < width*height; i++ {
		pixels[i*3] = bgColor[0]
		pixels[i*3+1] = bgColor[1]
		pixels[i*3+2] = bgColor[2]
	}

	// Calculate bar dimensions
	totalBarsWidth := numBars*barWidth + (numBars-1)*barSpacing
	startX := (width - totalBarsWidth) / 2

	// Draw each bar
	for i, value := range values {
		if i >= numBars {
			break
		}

		// Calculate bar height (0-100 maps to 0-50 pixels, leaving room for text)
		barHeight := int(value * 50.0 / 100.0)
		if barHeight < 1 {
			barHeight = 1
		}

		// Calculate x position for this bar
		x := startX + i*(barWidth+barSpacing)

		// Get color
		color := getBarColor(value)

		// Draw the bar (from bottom up)
		barBottom := 55 // Leave room at bottom
		for row := 0; row < barHeight; row++ {
			y := barBottom - row
			if y < 0 {
				break
			}

			// Fill the bar width
			for col := 0; col < barWidth; col++ {
				px := x + col
				if px >= 0 && px < width && y >= 0 && y < height {
					idx := (y*width + px) * 3
					pixels[idx] = color[0]
					pixels[idx+1] = color[1]
					pixels[idx+2] = color[2]
				}
			}
		}
	}

	return pixels
}

func main() {
	// IMPORTANT: Replace with your device IP
	deviceIP := "192.168.1.180"
	client := divoom.NewClient(deviceIP)

	fmt.Println("🚀 Bar Chart Display on PIXOO64")
	fmt.Println("   Device IP:", deviceIP)
	fmt.Println("   Simple vertical bars - easier to see!")
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

	fmt.Println("📊 Starting bar chart updates...")
	fmt.Println()

	// Keep rolling history
	history := []float64{}

	// Main loop - update every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	counter := 0.0
	picID := 1

	for range ticker.C {
		// Get new value
		value := simulateMetric(counter)
		counter++

		// Add to history (keep last numBars values)
		history = append(history, value)
		if len(history) > numBars {
			history = history[1:]
		}

		// Draw bar chart
		pixels := drawBarChart(history)

		// Create slightly different second frame
		pixels2 := make([]byte, len(pixels))
		copy(pixels2, pixels)
		for i := 0; i < len(pixels2); i++ {
			if pixels2[i] < 254 {
				pixels2[i]++
			}
		}

		// Reset every 10 frames
		if picID > 10 {
			client.ResetGifID()
			time.Sleep(500 * time.Millisecond)
			picID = 1
		}

		// Send frame 1
		frame1Data := base64.StdEncoding.EncodeToString(pixels)
		if err := client.SendGif(divoom.GifParams{
			PicNum:    2,
			PicWidth:  64,
			PicOffset: 0,
			PicID:     picID,
			PicSpeed:  500,
			PicData:   frame1Data,
		}); err != nil {
			log.Printf("Error: %v", err)
			continue
		}
		time.Sleep(100 * time.Millisecond)

		// Send frame 2
		frame2Data := base64.StdEncoding.EncodeToString(pixels2)
		if err := client.SendGif(divoom.GifParams{
			PicNum:    2,
			PicWidth:  64,
			PicOffset: 1,
			PicID:     picID,
			PicSpeed:  500,
			PicData:   frame2Data,
		}); err != nil {
			log.Printf("Error: %v", err)
			continue
		}
		time.Sleep(100 * time.Millisecond)

		picID++

		// Add text overlay with current value
		color := "#00FF00"
		if value >= 75 {
			color = "#FF0000"
		} else if value >= 50 {
			color = "#FFFF00"
		}

		if err := client.SendText(divoom.TextParams{
			TextID:     1,
			X:          0,
			Y:          2,
			Font:       4,
			TextWidth:  64,
			TextString: fmt.Sprintf("%.0f%%", value),
			Color:      color,
			Align:      2,
		}); err != nil {
			log.Printf("Error: %v", err)
		}

		// Console output
		fmt.Printf("📊 Value: %.1f%% | Bars: %d\n", value, len(history))
	}
}
