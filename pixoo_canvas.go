package divoom

import (
	"encoding/base64"
	"sync"
)

// PixooCanvas is a direct port of the Python library's approach
// Uses single-frame GIFs with incrementing PicID
type PixooCanvas struct {
	width   int
	height  int
	pixels  []byte
	client  *Client
	counter int
	mu      sync.Mutex

	// Auto-refresh settings (same as Python library)
	refreshCounterLimit         int
	refreshConnectionAutomatically bool
}

// NewPixooCanvas creates a canvas matching the Python library behavior
func NewPixooCanvas(client *Client) *PixooCanvas {
	return &PixooCanvas{
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
func (c *PixooCanvas) Clear(r, g, b byte) {
	for i := 0; i < c.width*c.height; i++ {
		c.pixels[i*3] = r
		c.pixels[i*3+1] = g
		c.pixels[i*3+2] = b
	}
}

// SetPixel sets a single pixel
func (c *PixooCanvas) SetPixel(x, y int, r, g, b byte) {
	if x < 0 || x >= c.width || y < 0 || y >= c.height {
		return
	}
	idx := (y*c.width + x) * 3
	c.pixels[idx] = r
	c.pixels[idx+1] = g
	c.pixels[idx+2] = b
}

// DrawLine draws a line (Bresenham's algorithm)
func (c *PixooCanvas) DrawLine(x1, y1, x2, y2 int, r, g, b byte) {
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
func (c *PixooCanvas) FillRectangle(x1, y1, x2, y2 int, r, g, b byte) {
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

// Push sends the canvas to device - EXACT port of Python library
func (c *PixooCanvas) Push() error {
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
		PicNum:    1,           // ← Single frame!
		PicWidth:  c.width,
		PicOffset: 0,
		PicID:     c.counter,   // ← Incrementing counter!
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
func (c *PixooCanvas) resetCounter() error {
	return c.client.ResetGifID()
}

// SetRefreshLimit sets the counter limit before reset (default: 32)
func (c *PixooCanvas) SetRefreshLimit(limit int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.refreshCounterLimit = limit
}
