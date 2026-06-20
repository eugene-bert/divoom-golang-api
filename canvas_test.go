package divoom

import (
	"image"
	"image/color"
	"sync"
	"testing"
)

func newTestCanvas() *Canvas {
	return &Canvas{
		width:                          64,
		height:                         64,
		pixels:                         make([]byte, 64*64*3),
		counter:                        1,
		refreshCounterLimit:            32,
		refreshConnectionAutomatically: true,
	}
}

func TestSetPixel(t *testing.T) {
	c := newTestCanvas()

	c.SetPixel(10, 20, 255, 128, 0)
	idx := (20*64 + 10) * 3
	if c.pixels[idx] != 255 || c.pixels[idx+1] != 128 || c.pixels[idx+2] != 0 {
		t.Errorf("SetPixel(10,20) = (%d,%d,%d), want (255,128,0)",
			c.pixels[idx], c.pixels[idx+1], c.pixels[idx+2])
	}
}

func TestSetPixelOutOfBounds(t *testing.T) {
	c := newTestCanvas()

	c.SetPixel(-1, 0, 255, 0, 0)
	c.SetPixel(0, -1, 255, 0, 0)
	c.SetPixel(64, 0, 255, 0, 0)
	c.SetPixel(0, 64, 255, 0, 0)

	for i, b := range c.pixels {
		if b != 0 {
			t.Fatalf("out-of-bounds SetPixel wrote to index %d", i)
		}
	}
}

func TestClear(t *testing.T) {
	c := newTestCanvas()

	c.Clear(100, 200, 50)

	for i := 0; i < 64*64; i++ {
		r, g, b := c.pixels[i*3], c.pixels[i*3+1], c.pixels[i*3+2]
		if r != 100 || g != 200 || b != 50 {
			t.Fatalf("Clear: pixel %d = (%d,%d,%d), want (100,200,50)", i, r, g, b)
		}
	}
}

func TestDrawLineHorizontal(t *testing.T) {
	c := newTestCanvas()

	c.DrawLine(5, 10, 15, 10, 255, 0, 0)

	for x := 5; x <= 15; x++ {
		idx := (10*64 + x) * 3
		if c.pixels[idx] != 255 {
			t.Errorf("horizontal line missing pixel at (%d, 10)", x)
		}
	}
}

func TestDrawLineVertical(t *testing.T) {
	c := newTestCanvas()

	c.DrawLine(10, 5, 10, 15, 0, 255, 0)

	for y := 5; y <= 15; y++ {
		idx := (y*64 + 10) * 3
		if c.pixels[idx+1] != 255 {
			t.Errorf("vertical line missing pixel at (10, %d)", y)
		}
	}
}

func TestDrawLineDiagonal(t *testing.T) {
	c := newTestCanvas()

	c.DrawLine(0, 0, 10, 10, 0, 0, 255)

	for i := 0; i <= 10; i++ {
		idx := (i*64 + i) * 3
		if c.pixels[idx+2] != 255 {
			t.Errorf("diagonal line missing pixel at (%d, %d)", i, i)
		}
	}
}

func TestFillRectangle(t *testing.T) {
	c := newTestCanvas()

	c.FillRectangle(10, 20, 30, 40, 255, 0, 0)

	for y := 20; y <= 40; y++ {
		for x := 10; x <= 30; x++ {
			idx := (y*64 + x) * 3
			if c.pixels[idx] != 255 {
				t.Fatalf("FillRectangle missing pixel at (%d, %d)", x, y)
			}
		}
	}

	idx := (19*64 + 10) * 3
	if c.pixels[idx] != 0 {
		t.Error("FillRectangle wrote outside bounds (above)")
	}
}

func TestFillRectangleSwapped(t *testing.T) {
	c := newTestCanvas()

	c.FillRectangle(30, 40, 10, 20, 255, 0, 0)

	idx := (30*64 + 20) * 3
	if c.pixels[idx] != 255 {
		t.Error("FillRectangle with swapped coords should normalize")
	}
}

func TestDrawCircleSymmetry(t *testing.T) {
	c := newTestCanvas()

	c.DrawCircle(32, 32, 10, 255, 0, 0)

	check := func(x, y int) bool {
		if x < 0 || x >= 64 || y < 0 || y >= 64 {
			return true
		}
		return c.pixels[(y*64+x)*3] == 255
	}

	for y := 0; y < 64; y++ {
		for x := 0; x < 64; x++ {
			if c.pixels[(y*64+x)*3] == 255 {
				dx, dy := x-32, y-32
				if !check(32-dx, 32+dy) || !check(32+dx, 32-dy) || !check(32-dx, 32-dy) {
					t.Fatalf("circle not symmetric at offset (%d, %d)", dx, dy)
				}
			}
		}
	}
}

func TestFillCircle(t *testing.T) {
	c := newTestCanvas()

	c.FillCircle(32, 32, 5, 255, 0, 0)

	filled := 0
	for y := 0; y < 64; y++ {
		for x := 0; x < 64; x++ {
			if c.pixels[(y*64+x)*3] == 255 {
				filled++
				dx, dy := x-32, y-32
				if dx*dx+dy*dy > 25 {
					t.Fatalf("FillCircle pixel (%d,%d) outside radius", x, y)
				}
			}
		}
	}

	if filled == 0 {
		t.Error("FillCircle drew nothing")
	}
}

func TestDrawImage(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	img.Set(0, 0, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	img.Set(1, 0, color.RGBA{R: 0, G: 255, B: 0, A: 255})
	img.Set(0, 1, color.RGBA{R: 0, G: 0, B: 255, A: 255})
	img.Set(1, 1, color.RGBA{R: 255, G: 255, B: 0, A: 255})

	c := newTestCanvas()
	c.DrawImage(img, 5, 5)

	check := func(x, y int, er, eg, eb byte) {
		idx := (y*64 + x) * 3
		if c.pixels[idx] != er || c.pixels[idx+1] != eg || c.pixels[idx+2] != eb {
			t.Errorf("DrawImage pixel (%d,%d) = (%d,%d,%d), want (%d,%d,%d)",
				x, y, c.pixels[idx], c.pixels[idx+1], c.pixels[idx+2], er, eg, eb)
		}
	}

	check(5, 5, 255, 0, 0)
	check(6, 5, 0, 255, 0)
	check(5, 6, 0, 0, 255)
	check(6, 6, 255, 255, 0)
}

func TestDrawImageResized(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 4, 4))
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			img.Set(x, y, color.RGBA{R: 255, G: 0, B: 0, A: 255})
		}
	}

	c := newTestCanvas()
	c.DrawImageResized(img, 0, 0, 64, 64)

	for y := 0; y < 64; y++ {
		for x := 0; x < 64; x++ {
			idx := (y*64 + x) * 3
			if c.pixels[idx] != 255 {
				t.Fatalf("DrawImageResized pixel (%d,%d) = %d, want 255", x, y, c.pixels[idx])
			}
		}
	}
}

func TestDrawImageFill(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 128, 128))
	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			img.Set(x, y, color.RGBA{R: 0, G: 128, B: 255, A: 255})
		}
	}

	c := newTestCanvas()
	c.DrawImageFill(img)

	idx := (32*64 + 32) * 3
	if c.pixels[idx] != 0 || c.pixels[idx+1] != 128 || c.pixels[idx+2] != 255 {
		t.Errorf("DrawImageFill center pixel = (%d,%d,%d), want (0,128,255)",
			c.pixels[idx], c.pixels[idx+1], c.pixels[idx+2])
	}
}

func TestConcurrentAccess(t *testing.T) {
	c := newTestCanvas()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(v byte) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				c.SetPixel(j%64, j%64, v, v, v)
				c.Clear(v, v, v)
				c.DrawLine(0, 0, 63, 63, v, v, v)
				c.FillRectangle(0, 0, 10, 10, v, v, v)
			}
		}(byte(i))
	}
	wg.Wait()
}
