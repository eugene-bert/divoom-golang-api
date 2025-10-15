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
	width  = 64
	height = 64
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

// drawLineGraph renders a line graph to 64x64 RGB pixels
func drawLineGraph(data []float64, min, max float64, lineColor, bgColor [3]byte) []byte {
	pixels := make([]byte, width*height*3)

	// Fill background
	for i := 0; i < width*height; i++ {
		pixels[i*3] = bgColor[0]
		pixels[i*3+1] = bgColor[1]
		pixels[i*3+2] = bgColor[2]
	}

	if len(data) < 2 {
		return pixels
	}

	// Draw line graph
	for i := 0; i < len(data)-1; i++ {
		// Map data point to x position
		x1 := i * width / len(data)
		x2 := (i + 1) * width / len(data)

		// Map data value to y position (inverted - 0 is top, 63 is bottom)
		y1 := height - 1 - int((data[i]-min)/(max-min)*float64(height-1))
		y2 := height - 1 - int((data[i+1]-min)/(max-min)*float64(height-1))

		// Clamp y values
		y1 = clamp(y1, 0, height-1)
		y2 = clamp(y2, 0, height-1)

		// Draw line from (x1,y1) to (x2,y2)
		drawLine(pixels, x1, y1, x2, y2, lineColor)
	}

	return pixels
}

// drawLine draws a line using Bresenham's algorithm
func drawLine(pixels []byte, x1, y1, x2, y2 int, color [3]byte) {
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	sx := 1
	if x1 > x2 {
		sx = -1
	}
	sy := 1
	if y1 > y2 {
		sy = -1
	}
	err := dx - dy

	for {
		setPixel(pixels, x1, y1, color)

		if x1 == x2 && y1 == y2 {
			break
		}

		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

// setPixel sets a pixel at (x, y) to the given color
func setPixel(pixels []byte, x, y int, color [3]byte) {
	if x < 0 || x >= width || y < 0 || y >= height {
		return
	}
	idx := (y*width + x) * 3
	pixels[idx] = color[0]
	pixels[idx+1] = color[1]
	pixels[idx+2] = color[2]
}

// clamp restricts value to [min, max]
func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// abs returns absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// simulateMetric simulates a metric value (you'd replace this with real data)
func simulateMetric(t float64) float64 {
	// Simulates CPU load: base 30% + sine wave + random noise
	base := 30.0
	wave := 30.0 * math.Sin(t/10.0)
	noise := rand.Float64()*20.0 - 10.0
	value := base + wave + noise

	// Clamp to 0-100
	if value < 0 {
		value = 0
	}
	if value > 100 {
		value = 100
	}
	return value
}

// getColor returns color based on value (green=low, yellow=medium, red=high)
func getColor(value float64) [3]byte {
	if value < 50 {
		return [3]byte{0, 255, 0} // Green
	} else if value < 75 {
		return [3]byte{255, 255, 0} // Yellow
	}
	return [3]byte{255, 0, 0} // Red
}

func main() {
	// IMPORTANT: Replace with your device IP
	deviceIP := "192.168.1.180"
	client := divoom.NewClient(deviceIP)

	fmt.Println("🚀 Starting Real-Time Graph Display on PIXOO64")
	fmt.Println("   Device IP:", deviceIP)
	fmt.Println("   Press Ctrl+C to stop")
	fmt.Println()

	// Create rolling data buffer (64 points = 64 pixels wide)
	data := NewMetricData(64)

	// Setup device
	fmt.Println("⚙️  Setting up device...")
	client.ResetGifID()
	time.Sleep(500 * time.Millisecond)
	client.SetChannelIndex(3)
	time.Sleep(500 * time.Millisecond)
	client.SetCustomPageIndex(1)
	time.Sleep(2 * time.Second)

	fmt.Println("📊 Starting dynamic graph updates...")
	fmt.Println()

	// Main loop - update every 2 seconds (slower to avoid overloading device)
	ticker := time.NewTicker(2 * time.Second)
	counter := 0.0
	picID := 1

	for range ticker.C {
		// 1. Get new metric value (simulate or get real data)
		value := simulateMetric(counter)
		data.Add(value)
		counter++

		// 2. Draw graph
		lineColor := getColor(value)
		bgColor := [3]byte{0, 0, 10} // Dark blue background
		pixels := drawLineGraph(data.GetValues(), 0, 100, lineColor, bgColor)

		// Create slightly different second frame (required by device)
		pixels2 := make([]byte, len(pixels))
		copy(pixels2, pixels)
		for i := 0; i < len(pixels2); i++ {
			if pixels2[i] < 255 {
				pixels2[i] = byte(clamp(int(pixels2[i])+1, 0, 255))
			}
		}

		// 3. Send to device
		frame1Data := base64.StdEncoding.EncodeToString(pixels)
		frame2Data := base64.StdEncoding.EncodeToString(pixels2)

		// Use incrementing PicID and reset every 10 frames to avoid accumulation
		if picID > 10 {
			client.ResetGifID()
			time.Sleep(500 * time.Millisecond)
			picID = 1
		}

		// Send frame 1
		if err := client.SendGif(divoom.GifParams{
			PicNum:    2,
			PicWidth:  64,
			PicOffset: 0,
			PicID:     picID,
			PicSpeed:  500,
			PicData:   frame1Data,
		}); err != nil {
			log.Printf("Error sending frame 1: %v", err)
			continue
		}
		time.Sleep(100 * time.Millisecond)

		// Send frame 2
		if err := client.SendGif(divoom.GifParams{
			PicNum:    2,
			PicWidth:  64,
			PicOffset: 1,
			PicID:     picID,
			PicSpeed:  500,
			PicData:   frame2Data,
		}); err != nil {
			log.Printf("Error sending frame 2: %v", err)
			continue
		}
		time.Sleep(100 * time.Millisecond)

		picID++

		// 4. Overlay current value as text
		colorStr := "#00FF00"
		if value >= 75 {
			colorStr = "#FF0000"
		} else if value >= 50 {
			colorStr = "#FFFF00"
		}

		if err := client.SendText(divoom.TextParams{
			TextID:     1,
			X:          0,
			Y:          2,
			Font:       2,
			TextWidth:  64,
			TextString: fmt.Sprintf("%.1f%%", value),
			Color:      colorStr,
			Align:      2,
		}); err != nil {
			log.Printf("Error sending text: %v", err)
		}

		// Print to console
		fmt.Printf("📈 Updated: %.1f%% | Data points: %d\n", value, len(data.GetValues()))
	}
}
