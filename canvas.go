package divoom

import (
	"encoding/base64"
	"time"
)

// Canvas represents a 64x64 pixel buffer for drawing
type Canvas struct {
	width  int
	height int
	pixels []byte // RGB data: width * height * 3 bytes
	client *Client
}

// NewCanvas creates a new drawing canvas
func NewCanvas(client *Client) *Canvas {
	return &Canvas{
		width:  64,
		height: 64,
		pixels: make([]byte, 64*64*3),
		client: client,
	}
}

// Clear fills the entire canvas with a color
func (c *Canvas) Clear(r, g, b byte) {
	for i := 0; i < c.width*c.height; i++ {
		c.pixels[i*3] = r
		c.pixels[i*3+1] = g
		c.pixels[i*3+2] = b
	}
}

// SetPixel sets a single pixel at (x, y) to the specified color
func (c *Canvas) SetPixel(x, y int, r, g, b byte) {
	if x < 0 || x >= c.width || y < 0 || y >= c.height {
		return
	}
	idx := (y*c.width + x) * 3
	c.pixels[idx] = r
	c.pixels[idx+1] = g
	c.pixels[idx+2] = b
}

// GetPixel returns the color of the pixel at (x, y)
func (c *Canvas) GetPixel(x, y int) (r, g, b byte) {
	if x < 0 || x >= c.width || y < 0 || y >= c.height {
		return 0, 0, 0
	}
	idx := (y*c.width + x) * 3
	return c.pixels[idx], c.pixels[idx+1], c.pixels[idx+2]
}

// DrawLine draws a line from (x1, y1) to (x2, y2) using Bresenham's algorithm
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

// DrawRectangle draws a rectangle outline
func (c *Canvas) DrawRectangle(x1, y1, x2, y2 int, r, g, b byte) {
	// Top line
	c.DrawLine(x1, y1, x2, y1, r, g, b)
	// Bottom line
	c.DrawLine(x1, y2, x2, y2, r, g, b)
	// Left line
	c.DrawLine(x1, y1, x1, y2, r, g, b)
	// Right line
	c.DrawLine(x2, y1, x2, y2, r, g, b)
}

// FillRectangle draws a filled rectangle
func (c *Canvas) FillRectangle(x1, y1, x2, y2 int, r, g, b byte) {
	// Ensure x1 <= x2 and y1 <= y2
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

// Push sends the canvas to the display
// WARNING: Don't call this more than once per second to avoid device issues
func (c *Canvas) Push() error {
	// Create slightly different second frame (required by device)
	pixels2 := make([]byte, len(c.pixels))
	copy(pixels2, c.pixels)
	for i := 0; i < len(pixels2); i++ {
		if pixels2[i] < 254 {
			pixels2[i]++
		}
	}

	// Encode frames
	frame1Data := base64.StdEncoding.EncodeToString(c.pixels)
	frame2Data := base64.StdEncoding.EncodeToString(pixels2)

	// Send frame 1
	if err := c.client.SendGif(GifParams{
		PicNum:    2,
		PicWidth:  64,
		PicOffset: 0,
		PicID:     1,
		PicSpeed:  500,
		PicData:   frame1Data,
	}); err != nil {
		return err
	}
	time.Sleep(100 * time.Millisecond)

	// Send frame 2
	if err := c.client.SendGif(GifParams{
		PicNum:    2,
		PicWidth:  64,
		PicOffset: 1,
		PicID:     1,
		PicSpeed:  500,
		PicData:   frame2Data,
	}); err != nil {
		return err
	}

	return nil
}

// abs returns absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
