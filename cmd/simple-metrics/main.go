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

// createBackgroundWithLabels creates a static background with labels
func createBackgroundWithLabels() []byte {
	width := 64
	height := 64
	pixels := make([]byte, width*height*3)

	// Very dark blue background (not pure black, so you can see it's working)
	bgColor := [3]byte{0, 0, 15}
	for i := 0; i < width*height; i++ {
		pixels[i*3] = bgColor[0]
		pixels[i*3+1] = bgColor[1]
		pixels[i*3+2] = bgColor[2]
	}

	return pixels
}

func main() {
	// IMPORTANT: Replace with your device IP
	deviceIP := "192.168.1.180"
	client := divoom.NewClient(deviceIP)

	fmt.Println("🚀 Simple Metrics Display")
	fmt.Println("   Just clean text, no BS")
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

	// Create background ONCE
	fmt.Println("🎨 Creating background...")
	pixels := createBackgroundWithLabels()
	pixels2 := make([]byte, len(pixels))
	copy(pixels2, pixels)
	for i := 0; i < len(pixels2); i++ {
		if pixels2[i] < 254 {
			pixels2[i]++
		}
	}

	frame1 := base64.StdEncoding.EncodeToString(pixels)
	frame2 := base64.StdEncoding.EncodeToString(pixels2)

	client.SendGif(divoom.GifParams{
		PicNum: 2, PicWidth: 64, PicOffset: 0,
		PicID: 1, PicSpeed: 500, PicData: frame1,
	})
	time.Sleep(100 * time.Millisecond)
	client.SendGif(divoom.GifParams{
		PicNum: 2, PicWidth: 64, PicOffset: 1,
		PicID: 1, PicSpeed: 500, PicData: frame2,
	})
	time.Sleep(500 * time.Millisecond)

	fmt.Println("📊 Running (updating all text each cycle)...\n")

	// Main loop - update every 1 second
	ticker := time.NewTicker(1 * time.Second)
	counter := 0.0

	for range ticker.C {
		// Get metrics
		cpuLoad := simulateMetric(counter)
		memLoad := simulateMetric(counter + 3)
		diskLoad := simulateMetric(counter + 7)
		counter++

		// Send ALL text each cycle (labels + values together)
		// This ensures everything stays visible

		const labelColor = "#AAAAAA"
		const formatStr = "%3.0f%%"

		// CPU label
		client.SendText(divoom.TextParams{
			TextID: 1, X: 2, Y: 12, Font: 2,
			TextWidth: 25, TextString: "CPU",
			Color: labelColor, Align: 1,
		})
		// CPU value
		cpuColor := getColor(cpuLoad)
		client.SendText(divoom.TextParams{
			TextID: 2, X: 28, Y: 10, Font: 5,
			TextWidth: 34, TextString: fmt.Sprintf(formatStr, cpuLoad),
			Color: cpuColor, Align: 3,
		})

		// MEM label
		client.SendText(divoom.TextParams{
			TextID: 3, X: 2, Y: 30, Font: 2,
			TextWidth: 25, TextString: "MEM",
			Color: labelColor, Align: 1,
		})
		// Memory value
		memColor := getColor(memLoad)
		client.SendText(divoom.TextParams{
			TextID: 4, X: 28, Y: 28, Font: 5,
			TextWidth: 34, TextString: fmt.Sprintf(formatStr, memLoad),
			Color: memColor, Align: 3,
		})

		// DSK label
		client.SendText(divoom.TextParams{
			TextID: 5, X: 2, Y: 48, Font: 2,
			TextWidth: 25, TextString: "DSK",
			Color: labelColor, Align: 1,
		})
		// Disk value
		diskColor := getColor(diskLoad)
		client.SendText(divoom.TextParams{
			TextID: 6, X: 28, Y: 46, Font: 5,
			TextWidth: 34, TextString: fmt.Sprintf(formatStr, diskLoad),
			Color: diskColor, Align: 3,
		})

		// Console
		fmt.Printf("CPU: %5.1f%% | MEM: %5.1f%% | DSK: %5.1f%%\n",
			cpuLoad, memLoad, diskLoad)
	}
}
