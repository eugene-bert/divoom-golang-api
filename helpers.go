package divoom

import (
	"encoding/base64"
	"time"
)

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// SendBlankScreen sends a blank black screen to the device on Custom channel page 1
// IMPORTANT: You must be on Custom channel (3) and CustomPageIndex 1 for this to display
// Note: Sends a 2-frame animation with DIFFERENT frames (device requires this)
func (c *Client) SendBlankScreen() error {
	// Frame 0: Pure black (0,0,0)
	frame0 := make([]byte, 64*64*3)
	// All bytes are 0 (black)

	// Frame 1: Very dark gray (2,2,2 for ALL pixels)
	// Must be different from frame 0 for device to recognize as animation
	frame1 := make([]byte, 64*64*3)
	for i := 0; i < 64*64*3; i++ {
		frame1[i] = 2 // Very slightly gray (imperceptible but different)
	}

	frame0Data := base64.StdEncoding.EncodeToString(frame0)
	frame1Data := base64.StdEncoding.EncodeToString(frame1)

	// Send frame 0 (pure black)
	// Use PicID 1 (we reset GIF ID before calling this)
	if err := c.SendGif(GifParams{
		PicNum:    2,
		PicWidth:  64,
		PicOffset: 0,
		PicID:     1, // Start from 1 after reset
		PicSpeed:  500,
		PicData:   frame0Data,
	}); err != nil {
		return err
	}

	time.Sleep(200 * time.Millisecond)

	// Send frame 1 (dark gray)
	return c.SendGif(GifParams{
		PicNum:    2,
		PicWidth:  64,
		PicOffset: 1,
		PicID:     1, // Same PicID for both frames
		PicSpeed:  500,
		PicData:   frame1Data,
	})
}

// SendColorScreen sends a solid color screen to the device on Custom channel page 1
// IMPORTANT: You must be on Custom channel (3) and CustomPageIndex 1 for this to display
// color should be in RGB format (e.g., 0x000000 for black, 0xFFFFFF for white)
// Note: Sends a 2-frame animation (both same color) as single-frame GIFs may not display
func (c *Client) SendColorScreen(rgb uint32) error {
	r := byte((rgb >> 16) & 0xFF)
	g := byte((rgb >> 8) & 0xFF)
	b := byte(rgb & 0xFF)

	// Create a 64x64 screen with the specified color
	pixels := make([]byte, 64*64*3)
	for i := 0; i < 64*64; i++ {
		pixels[i*3] = r
		pixels[i*3+1] = g
		pixels[i*3+2] = b
	}

	// Encode to base64
	base64Data := base64.StdEncoding.EncodeToString(pixels)

	// Create slightly different frame 1 (add 1 to all RGB values)
	pixels1 := make([]byte, 64*64*3)
	for i := 0; i < 64*64; i++ {
		pixels1[i*3] = byte(min(int(r)+1, 255))
		pixels1[i*3+1] = byte(min(int(g)+1, 255))
		pixels1[i*3+2] = byte(min(int(b)+1, 255))
	}
	base64Data1 := base64.StdEncoding.EncodeToString(pixels1)

	// Send frame 0
	if err := c.SendGif(GifParams{
		PicNum:    2,
		PicWidth:  64,
		PicOffset: 0,
		PicID:     1, // Use PicID 1 (after reset)
		PicSpeed:  500,
		PicData:   base64Data,
	}); err != nil {
		return err
	}

	time.Sleep(200 * time.Millisecond)

	// Send frame 1 (slightly different color)
	return c.SendGif(GifParams{
		PicNum:    2,
		PicWidth:  64,
		PicOffset: 1,
		PicID:     1, // Same PicID
		PicSpeed:  500,
		PicData:   base64Data1,
	})
}

// DisplayText is a convenience method that properly sets up the custom channel and displays text
// This is the recommended way to display text on the PIXOO64
// It automatically switches to Custom channel page 1, sends a blank screen, and overlays text
func (c *Client) DisplayText(text string, color string, opts ...TextOption) error {
	// Default options
	params := TextParams{
		TextID:     1,
		X:          0,
		Y:          24,
		Direction:  0,
		Font:       4,
		TextWidth:  64,
		Speed:      0,
		TextString: text,
		Color:      color,
		Align:      2, // center
	}

	// Apply options
	for _, opt := range opts {
		opt(&params)
	}

	// Step 1: Reset GIF ID to clear any accumulated GIF data
	if err := c.ResetGifID(); err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)

	// Step 2: Switch to Custom channel (required for custom content)
	if err := c.SetChannelIndex(3); err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)

	// Step 3: Switch to CustomPageIndex 1 (page 0 shows favorites, page 1 is for custom content)
	if err := c.SetCustomPageIndex(1); err != nil {
		return err
	}
	time.Sleep(2 * time.Second) // 2 seconds to ensure page switch completes

	// Step 4: Send a blank screen (creates the base animation)
	if err := c.SendBlankScreen(); err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond) // Wait for animation to be ready

	// Step 5: Overlay the text on the blank screen
	return c.SendText(params)
}
