package divoom

import (
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"sync"
)

// Canvas provides a drawing buffer for the Pixoo64 display
// Uses single-frame GIFs with incrementing PicID (based on Python library approach)
type Canvas struct {
	width   int
	height  int
	pixels  []byte
	client  *Client
	counter int
	mu      sync.Mutex

	// Auto-refresh settings (same as Python library)
	refreshCounterLimit            int
	refreshConnectionAutomatically bool
}

// NewCanvas creates a new drawing canvas
func NewCanvas(client *Client) *Canvas {
	return &Canvas{
		width:                          64,
		height:                         64,
		pixels:                         make([]byte, 64*64*3),
		client:                         client,
		counter:                        1,
		refreshCounterLimit:            32, // Same as Python
		refreshConnectionAutomatically: true,
	}
}

// Clear fills the canvas with a color
func (c *Canvas) Clear(r, g, b byte) {
	for i := 0; i < c.width*c.height; i++ {
		c.pixels[i*3] = r
		c.pixels[i*3+1] = g
		c.pixels[i*3+2] = b
	}
}

// SetPixel sets a single pixel
func (c *Canvas) SetPixel(x, y int, r, g, b byte) {
	if x < 0 || x >= c.width || y < 0 || y >= c.height {
		return
	}
	idx := (y*c.width + x) * 3
	c.pixels[idx] = r
	c.pixels[idx+1] = g
	c.pixels[idx+2] = b
}

// DrawLine draws a line (Bresenham's algorithm)
func (c *Canvas) DrawLine(x1, y1, x2, y2 int, r, g, b byte) {
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
		c.SetPixel(x1, y1, r, g, b)
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

// FillRectangle draws a filled rectangle
func (c *Canvas) FillRectangle(x1, y1, x2, y2 int, r, g, b byte) {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			c.SetPixel(x, y, r, g, b)
		}
	}
}

// DrawCircle draws a circle outline using midpoint circle algorithm
func (c *Canvas) DrawCircle(centerX, centerY, radius int, r, g, b byte) {
	x := radius
	y := 0
	err := 0

	for x >= y {
		c.SetPixel(centerX+x, centerY+y, r, g, b)
		c.SetPixel(centerX+y, centerY+x, r, g, b)
		c.SetPixel(centerX-y, centerY+x, r, g, b)
		c.SetPixel(centerX-x, centerY+y, r, g, b)
		c.SetPixel(centerX-x, centerY-y, r, g, b)
		c.SetPixel(centerX-y, centerY-x, r, g, b)
		c.SetPixel(centerX+y, centerY-x, r, g, b)
		c.SetPixel(centerX+x, centerY-y, r, g, b)

		if err <= 0 {
			y++
			err += 2*y + 1
		}
		if err > 0 {
			x--
			err -= 2*x + 1
		}
	}
}

// FillCircle draws a filled circle
func (c *Canvas) FillCircle(centerX, centerY, radius int, r, g, b byte) {
	for y := -radius; y <= radius; y++ {
		for x := -radius; x <= radius; x++ {
			if x*x+y*y <= radius*radius {
				c.SetPixel(centerX+x, centerY+y, r, g, b)
			}
		}
	}
}

// Push sends the canvas to device - EXACT port of Python library
func (c *Canvas) Push() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Reset counter if needed (every 32 frames, like Python)
	if c.refreshConnectionAutomatically && c.counter >= c.refreshCounterLimit {
		if err := c.resetCounter(); err != nil {
			return err
		}
		c.counter = 1
	}

	// Encode pixel data
	picData := base64.StdEncoding.EncodeToString(c.pixels)

	// Send SINGLE frame with incrementing PicID (EXACT Python behavior)
	err := c.client.SendGif(GifParams{
		PicNum:    1, // ← Single frame!
		PicWidth:  c.width,
		PicOffset: 0,
		PicID:     c.counter, // ← Incrementing counter!
		PicSpeed:  1000,
		PicData:   picData,
	})

	if err != nil {
		return err
	}

	// Increment counter for next push
	c.counter++

	return nil
}

// resetCounter resets the GIF ID counter (called every 32 frames)
func (c *Canvas) resetCounter() error {
	return c.client.ResetGifID()
}

// SetRefreshLimit sets the counter limit before reset (default: 32)
func (c *Canvas) SetRefreshLimit(limit int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.refreshCounterLimit = limit
}

// DrawImage draws a Go image.Image onto the canvas at position (x, y)
// The image is automatically resized to fit if it's larger than 64x64
func (c *Canvas) DrawImage(img image.Image, x, y int) {
	bounds := img.Bounds()
	for py := bounds.Min.Y; py < bounds.Min.Y+bounds.Dy(); py++ {
		for px := bounds.Min.X; px < bounds.Min.X+bounds.Dx(); px++ {
			targetX := x + (px - bounds.Min.X)
			targetY := y + (py - bounds.Min.Y)

			if targetX >= 0 && targetX < c.width && targetY >= 0 && targetY < c.height {
				r, g, b, _ := img.At(px, py).RGBA()
				// Convert from 16-bit to 8-bit
				c.SetPixel(targetX, targetY, byte(r>>8), byte(g>>8), byte(b>>8))
			}
		}
	}
}

// LoadImageFromFile loads an image file and returns it as image.Image
// Supports PNG, JPEG, and GIF formats
func LoadImageFromFile(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	return img, nil
}

// abs returns absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
