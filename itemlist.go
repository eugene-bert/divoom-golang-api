package divoom

// ItemListParams represents parameters for sending display item list
type ItemListParams struct {
	TextID     int    `json:"TextId"`
	Type       int    `json:"type"` // Display type (22 = custom text)
	X          int    `json:"x"`
	Y          int    `json:"y"`
	Direction  int    `json:"dir"`
	Font       int    `json:"font"`
	TextWidth  int    `json:"TextWidth"`
	TextHeight int    `json:"Textheight"`
	Speed      int    `json:"speed"`
	Align      int    `json:"align"`
	TextString string `json:"TextString,omitempty"`
	Color      string `json:"color"`
	UpdateTime int    `json:"update_time,omitempty"`
}

// SendHttpItemList sends a display item list (more reliable for text)
func (c *Client) SendHttpItemList(items []ItemListParams) error {
	payload := map[string]interface{}{
		"Command":  "Draw/SendHttpItemList",
		"ItemList": items,
	}
	_, err := c.sendCommand(payload)
	return err
}

// DisplayTextViaItemList displays text using the ItemList method
// This appears to be more reliable than SendText
func (c *Client) DisplayTextViaItemList(text string, color string) error {
	// First send blank screen
	if err := c.SendBlankScreen(); err != nil {
		return err
	}

	// Then send text as item list
	items := []ItemListParams{
		{
			TextID:     1,
			Type:       22, // TEXT_MESSAGE type
			X:          0,
			Y:          24,
			Direction:  0,
			Font:       4,
			TextWidth:  64,
			TextHeight: 16,
			Speed:      0,
			Align:      2,
			TextString: text,
			Color:      color,
		},
	}

	return c.SendHttpItemList(items)
}
