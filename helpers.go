package divoom

import (
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"time"
)

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

// DisplayText sets up the custom channel and displays text on the PIXOO64.
// It automatically switches to Custom channel page 1, sends a blank screen, and overlays text.
func (c *Client) DisplayText(text string, color string, opts ...TextOption) error {
	return c.DisplayTextContext(context.Background(), text, color, opts...)
}

// DisplayTextContext is like DisplayText but accepts a context for cancellation.
func (c *Client) DisplayTextContext(ctx context.Context, text string, color string, opts ...TextOption) error {
	params := TextParams{
		TextID: 1, X: 0, Y: 24, Direction: 0, Font: 4,
		TextWidth: 64, Speed: 0, TextString: text, Color: color, Align: 2,
	}
	for _, opt := range opts {
		opt(&params)
	}

	steps := []func() error{
		func() error { return c.ResetGifID() },
		func() error { return sleep(ctx, 500*time.Millisecond) },
		func() error { return c.SetChannelIndex(3) },
		func() error { return sleep(ctx, 500*time.Millisecond) },
		func() error { return c.SetCustomPageIndex(1) },
		func() error { return sleep(ctx, 2*time.Second) },
		func() error { return c.SendBlankScreen() },
		func() error { return sleep(ctx, 500*time.Millisecond) },
		func() error { return c.SendText(params) },
	}
	for _, step := range steps {
		if err := step(); err != nil {
			return err
		}
	}
	return nil
}

// PlayLocalGif loads and plays a GIF file from the local filesystem.
func (c *Client) PlayLocalGif(filePath string) error {
	return c.PlayLocalGifContext(context.Background(), filePath)
}

// PlayLocalGifContext is like PlayLocalGif but accepts a context for cancellation.
func (c *Client) PlayLocalGifContext(ctx context.Context, filePath string) error {
	// Open the GIF file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open GIF file: %w", err)
	}
	defer file.Close()

	// Decode the GIF
	gifImg, err := gif.DecodeAll(file)
	if err != nil {
		return fmt.Errorf("failed to decode GIF: %w", err)
	}

	if len(gifImg.Image) == 0 {
		return fmt.Errorf("GIF has no frames")
	}

	if err := c.ResetGifID(); err != nil {
		return err
	}
	if err := sleep(ctx, 100*time.Millisecond); err != nil {
		return err
	}

	for i, frame := range gifImg.Image {
		// Convert frame to 64x64 RGB bytes
		pixels := make([]byte, 64*64*3)
		bounds := frame.Bounds()

		for y := 0; y < 64; y++ {
			for x := 0; x < 64; x++ {
				// Scale coordinates if image is not 64x64
				srcX := bounds.Min.X + (x * bounds.Dx() / 64)
				srcY := bounds.Min.Y + (y * bounds.Dy() / 64)

				var col color.Color
				if srcX < bounds.Max.X && srcY < bounds.Max.Y {
					col = frame.At(srcX, srcY)
				} else {
					col = color.Black
				}

				r, g, b, _ := col.RGBA()
				idx := (y*64 + x) * 3
				pixels[idx] = byte(r >> 8)
				pixels[idx+1] = byte(g >> 8)
				pixels[idx+2] = byte(b >> 8)
			}
		}

		// Encode and send frame
		picData := base64.StdEncoding.EncodeToString(pixels)

		picSpeed := 100
		if i < len(gifImg.Delay) {
			picSpeed = gifImg.Delay[i] * 10
		}

		if err := c.SendGif(GifParams{
			PicNum:    len(gifImg.Image),
			PicWidth:  64,
			PicOffset: i,
			PicID:     1,
			PicSpeed:  picSpeed,
			PicData:   picData,
		}); err != nil {
			return fmt.Errorf("failed to send frame %d: %w", i, err)
		}

		if err := sleep(ctx, 100*time.Millisecond); err != nil {
			return err
		}
	}

	return nil
}

// DisplayImageFile loads an image file, resizes it to 64x64, and displays it.
// Supports PNG, JPEG, and GIF formats.
func (c *Client) DisplayImageFile(path string) error {
	return c.DisplayImageFileContext(context.Background(), path)
}

// DisplayImageFileContext is like DisplayImageFile but accepts a context for cancellation.
func (c *Client) DisplayImageFileContext(ctx context.Context, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	return c.displayImage(ctx, img)
}

// DisplayImageURL fetches an image from a URL, resizes it to 64x64, and displays it.
func (c *Client) DisplayImageURL(url string) error {
	return c.DisplayImageURLContext(context.Background(), url)
}

// DisplayImageURLContext is like DisplayImageURL but accepts a context for cancellation.
func (c *Client) DisplayImageURLContext(ctx context.Context, url string) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	imageClient := &http.Client{Timeout: 30 * time.Second}
	resp, err := imageClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to fetch image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected HTTP status: %d", resp.StatusCode)
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	return c.displayImage(ctx, img)
}

func (c *Client) displayImage(ctx context.Context, img image.Image) error {
	steps := []func() error{
		func() error { return c.ResetGifID() },
		func() error { return sleep(ctx, 500*time.Millisecond) },
		func() error { return c.SetChannelIndex(3) },
		func() error { return sleep(ctx, 500*time.Millisecond) },
		func() error { return c.SetCustomPageIndex(1) },
		func() error { return sleep(ctx, 2*time.Second) },
		func() error {
			canvas := NewCanvas(c)
			canvas.DrawImageFill(img)
			return canvas.Push()
		},
	}
	for _, step := range steps {
		if err := step(); err != nil {
			return err
		}
	}
	return nil
}

// sleep waits for duration or until context is cancelled.
func sleep(ctx context.Context, d time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(d):
		return nil
	}
}
