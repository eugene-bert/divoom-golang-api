package divoom

import "strings"

// TextOption is a function that modifies TextParams
type TextOption func(*TextParams)

// WithPosition sets the X and Y position
func WithPosition(x, y int) TextOption {
	return func(p *TextParams) {
		p.X = x
		p.Y = y
	}
}

// WithFont sets the font ID
func WithFont(font int) TextOption {
	return func(p *TextParams) {
		p.Font = font
	}
}

// WithAlignment sets the text alignment (1=left, 2=center, 3=right)
func WithAlignment(align int) TextOption {
	return func(p *TextParams) {
		p.Align = align
	}
}

// WithScroll enables scrolling with the specified speed in milliseconds
func WithScroll(speed int, direction int) TextOption {
	return func(p *TextParams) {
		p.Speed = speed
		p.Direction = direction
	}
}

// WithTextWidth sets the text width
func WithTextWidth(width int) TextOption {
	return func(p *TextParams) {
		p.TextWidth = width
	}
}

// ParseHexColor converts a hex color string to uppercase and ensures it has the # prefix
func ParseHexColor(color string) string {
	color = strings.ToUpper(strings.TrimSpace(color))
	if !strings.HasPrefix(color, "#") {
		color = "#" + color
	}
	return color
}
