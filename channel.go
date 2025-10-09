package divoom

import "fmt"

// SetBrightness sets the screen brightness
// brightness: 0-100
func (c *Client) SetBrightness(brightness int) error {
	if brightness < 0 || brightness > 100 {
		return fmt.Errorf("invalid brightness: %d (must be 0-100)", brightness)
	}

	payload := map[string]interface{}{
		"Command":    "Channel/SetBrightness",
		"Brightness": brightness,
	}
	_, err := c.sendCommand(payload)
	return err
}

// GetAllConf gets all device configuration settings
func (c *Client) GetAllConf() (*AllConfigResponse, error) {
	payload := map[string]string{
		"Command": "Channel/GetAllConf",
	}

	var result AllConfigResponse
	err := c.sendCommandWithResponse(payload, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrorCode != 0 {
		return nil, fmt.Errorf("device returned error code: %d", result.ErrorCode)
	}

	return &result, nil
}

// SetSubscribeGalleryTime sets the gallery subscription time
// time: time in seconds
func (c *Client) SetSubscribeGalleryTime(timeSeconds int) error {
	payload := map[string]interface{}{
		"Command": "Channel/SetSubscribeGalleryTime",
		"Time":    timeSeconds,
	}
	_, err := c.sendCommand(payload)
	return err
}
