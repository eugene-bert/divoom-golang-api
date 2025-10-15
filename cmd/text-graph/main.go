package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/eugene-bert/divoom-golang-api"
)

const (
	graphWidth  = 15 // Number of data points to show
	graphHeight = 6  // Height in character rows (reduced to fit better)
)

// MetricData holds rolling time-series data
type MetricData struct {
	values []float64
	maxLen int
}

// NewMetricData creates a new rolling data buffer
func NewMetricData(maxLen int) *MetricData {
	return &MetricData{
		values: make([]float64, 0, maxLen),
		maxLen: maxLen,
	}
}

// Add adds a new value and removes oldest if buffer is full
func (m *MetricData) Add(value float64) {
	m.values = append(m.values, value)
	if len(m.values) > m.maxLen {
		m.values = m.values[1:] // Remove oldest
	}
}

// GetValues returns current values
func (m *MetricData) GetValues() []float64 {
	return m.values
}

// Latest returns the most recent value
func (m *MetricData) Latest() float64 {
	if len(m.values) == 0 {
		return 0
	}
	return m.values[len(m.values)-1]
}

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

// createTextGraph creates a text-based graph using simple ASCII characters
func createTextGraph(data []float64, min, max float64) []string {
	lines := make([]string, graphHeight)

	if len(data) == 0 {
		for i := 0; i < graphHeight; i++ {
			lines[i] = ""
			for j := 0; j < graphWidth; j++ {
				lines[i] += "."
			}
		}
		return lines
	}

	// Initialize with background (dots)
	for i := 0; i < graphHeight; i++ {
		lines[i] = ""
		for j := 0; j < graphWidth; j++ {
			lines[i] += "."
		}
	}

	// Convert data values to character positions
	for i, value := range data {
		// Map data point to x position
		x := i * graphWidth / len(data)
		if x >= graphWidth {
			x = graphWidth - 1
		}

		// Map value to y position (inverted - 0 is top)
		normalizedValue := (value - min) / (max - min)
		yFloat := (1.0 - normalizedValue) * float64(graphHeight-1)
		y := int(yFloat)

		if y < 0 {
			y = 0
		}
		if y >= graphHeight {
			y = graphHeight - 1
		}

		// Fill from bottom to this height with # symbols
		for row := y; row < graphHeight; row++ {
			lineRunes := []rune(lines[row])
			lineRunes[x] = '#'
			lines[row] = string(lineRunes)
		}
	}

	return lines
}

func main() {
	// IMPORTANT: Replace with your device IP
	deviceIP := "192.168.1.180"
	client := divoom.NewClient(deviceIP)

	fmt.Println("🚀 Text-Based Real-Time Graph on PIXOO64")
	fmt.Println("   Device IP:", deviceIP)
	fmt.Println("   NO loading animations - pure text updates!")
	fmt.Println("   Press Ctrl+C to stop")
	fmt.Println()

	// Create rolling data buffer
	data := NewMetricData(graphWidth)

	// Setup device ONCE
	fmt.Println("⚙️  Setting up device...")
	client.ResetGifID()
	time.Sleep(500 * time.Millisecond)
	client.SetChannelIndex(3)
	time.Sleep(500 * time.Millisecond)
	client.SetCustomPageIndex(1)
	time.Sleep(2 * time.Second)

	// Send a simple dark background ONCE
	fmt.Println("🎨 Setting up background...")
	client.SendColorScreen(0x000008) // Very dark blue
	time.Sleep(500 * time.Millisecond)

	fmt.Println("📊 Starting real-time graph updates...")
	fmt.Println()

	// Main loop - update every 1 second
	ticker := time.NewTicker(1 * time.Second)
	counter := 0.0
	textID := 1

	for range ticker.C {
		// 1. Get new metric value
		value := simulateMetric(counter)
		data.Add(value)
		counter++

		// 2. Clear previous text
		if err := client.ClearText(); err != nil {
			log.Printf("Error clearing text: %v", err)
		}
		time.Sleep(50 * time.Millisecond)

		// Reset text ID counter
		textID = 1

		// 3. Display current value at top
		valueColor := getColor(value)
		if err := client.SendText(divoom.TextParams{
			TextID:     textID,
			X:          2,
			Y:          2,
			Font:       2,
			TextWidth:  60,
			TextString: fmt.Sprintf("%.1f%%", value),
			Color:      valueColor,
			Align:      1,
		}); err != nil {
			log.Printf("Error: %v", err)
		}
		textID++

		// 4. Create and display text-based graph
		graphLines := createTextGraph(data.GetValues(), 0, 100)

		startY := 10 // Start Y position for graph
		for i, line := range graphLines {
			if err := client.SendText(divoom.TextParams{
				TextID:     textID,
				X:          0,
				Y:          startY + (i * 6), // Space lines vertically
				Font:       1, // Slightly larger font for visibility
				TextWidth:  64,
				TextString: line,
				Color:      valueColor,
				Align:      1,
			}); err != nil {
				log.Printf("Error sending graph line %d: %v", i, err)
			}
			textID++

			// Debug output
			fmt.Printf("   Line %d (Y=%d): %s\n", i, startY+(i*6), line)
		}

		// 5. Display timestamp at bottom (fix Y position to be visible)
		if err := client.SendText(divoom.TextParams{
			TextID:     textID,
			X:          2,
			Y:          54,
			Font:       1,
			TextWidth:  60,
			TextString: time.Now().Format("15:04"),
			Color:      "#606060",
			Align:      1,
		}); err != nil {
			log.Printf("Error: %v", err)
		}

		// Console output
		fmt.Printf("📈 Value: %5.1f%% | Points: %d | Color: %s\n",
			value, len(data.GetValues()), valueColor)
	}
}
