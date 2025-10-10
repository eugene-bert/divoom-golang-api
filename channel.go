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

// SetChannelIndex sets the active channel
// selectIndex: 0=Faces, 1=Cloud Channel, 2=Visualizer, 3=Custom, 4=Black screen
func (c *Client) SetChannelIndex(selectIndex int) error {
	if selectIndex < 0 || selectIndex > 4 {
		return fmt.Errorf("invalid channel index: %d (must be 0-4)", selectIndex)
	}

	payload := map[string]interface{}{
		"Command":     "Channel/SetIndex",
		"SelectIndex": selectIndex,
	}
	_, err := c.sendCommand(payload)
	return err
}

// GetChannelIndex gets the current active channel index
func (c *Client) GetChannelIndex() (int, error) {
	payload := map[string]string{
		"Command": "Channel/GetIndex",
	}

	var result struct {
		ErrorCode   int `json:"error_code"`
		SelectIndex int `json:"SelectIndex"`
	}

	err := c.sendCommandWithResponse(payload, &result)
	if err != nil {
		return 0, err
	}

	if result.ErrorCode != 0 {
		return 0, fmt.Errorf("device returned error code: %d", result.ErrorCode)
	}

	return result.SelectIndex, nil
}

// SetCustomPageIndex sets the custom page index (0-2)
// Only works when channel is set to Custom (index 3)
func (c *Client) SetCustomPageIndex(customPageIndex int) error {
	if customPageIndex < 0 || customPageIndex > 2 {
		return fmt.Errorf("invalid custom page index: %d (must be 0-2)", customPageIndex)
	}

	payload := map[string]interface{}{
		"Command":         "Channel/SetCustomPageIndex",
		"CustomPageIndex": customPageIndex,
	}
	_, err := c.sendCommand(payload)
	return err
}
