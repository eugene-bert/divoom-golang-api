package divoom

// SendText sends text to the display
func (c *Client) SendText(params TextParams) error {
	payload := map[string]interface{}{
		"Command":    "Draw/SendHttpText",
		"TextId":     params.TextID,
		"x":          params.X,
		"y":          params.Y,
		"dir":        params.Direction,
		"font":       params.Font,
		"TextWidth":  params.TextWidth,
		"speed":      params.Speed,
		"TextString": params.TextString,
		"color":      params.Color,
		"align":      params.Align,
	}
	_, err := c.sendCommand(payload)
	return err
}

// ClearText clears all text areas from the display
func (c *Client) ClearText() error {
	payload := map[string]string{
		"Command": "Draw/ClearHttpText",
	}
	_, err := c.sendCommand(payload)
	return err
}

// SendGif sends a GIF animation to the display
func (c *Client) SendGif(params GifParams) error {
	payload := map[string]interface{}{
		"Command":   "Draw/SendHttpGif",
		"PicNum":    params.PicNum,
		"PicWidth":  params.PicWidth,
		"PicOffset": params.PicOffset,
		"PicID":     params.PicID,
		"PicSpeed":  params.PicSpeed,
		"PicData":   params.PicData,
	}
	_, err := c.sendCommand(payload)
	return err
}

// ResetGifID resets the GIF ID counter
func (c *Client) ResetGifID() error {
	payload := map[string]string{
		"Command": "Draw/ResetHttpGifId",
	}
	_, err := c.sendCommand(payload)
	return err
}

// SendRemote sends a remote control command
// key: the remote key code
func (c *Client) SendRemote(key int) error {
	payload := map[string]interface{}{
		"Command": "Draw/SendRemote",
		"Key":     key,
	}
	_, err := c.sendCommand(payload)
	return err
}

// UseHTTPCommandSource enables/disables HTTP command source
func (c *Client) UseHTTPCommandSource(enabled bool) error {
	flag := 0
	if enabled {
		flag = 1
	}

	payload := map[string]interface{}{
		"Command": "Draw/UseHTTPCommandSource",
		"flag":    flag,
	}
	_, err := c.sendCommand(payload)
	return err
}
