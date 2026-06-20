package divoom

import "fmt"

// SysReboot reboots the device
func (c *Client) SysReboot() error {
	payload := map[string]string{
		"Command": "Device/SysReboot",
	}
	_, err := c.sendCommand(payload)
	return err
}

// GetDeviceTime gets the current device time
func (c *Client) GetDeviceTime() (*DeviceTime, error) {
	payload := map[string]string{
		"Command": "Device/GetDeviceTime",
	}

	var result DeviceTime
	if err := c.sendCommandWithResponse(payload, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SetUTC sets the device UTC time
func (c *Client) SetUTC(utc int64) error {
	payload := map[string]interface{}{
		"Command": "Device/SetUTC",
		"Utc":     utc,
	}
	_, err := c.sendCommand(payload)
	return err
}

// OnOffScreen turns the screen on or off
func (c *Client) OnOffScreen(on bool) error {
	onOff := 0
	if on {
		onOff = 1
	}

	payload := map[string]interface{}{
		"Command": "Channel/OnOffScreen",
		"OnOff":   onOff,
	}
	_, err := c.sendCommand(payload)
	return err
}

// SetDisTempMode sets the temperature display mode
// mode: 0 = Celsius, 1 = Fahrenheit
func (c *Client) SetDisTempMode(mode int) error {
	payload := map[string]interface{}{
		"Command": "Device/SetDisTempMode",
		"Mode":    mode,
	}
	_, err := c.sendCommand(payload)
	return err
}

// SetScreenRotationAngle sets the screen rotation angle
// angle: 0 = normal, 1 = 90°, 2 = 180°, 3 = 270°
func (c *Client) SetScreenRotationAngle(angle int) error {
	if angle < 0 || angle > 3 {
		return fmt.Errorf("invalid angle: %d (must be 0-3)", angle)
	}

	payload := map[string]interface{}{
		"Command": "Device/SetScreenRotationAngle",
		"Mode":    angle,
	}
	_, err := c.sendCommand(payload)
	return err
}

// SetMirrorMode enables or disables mirror mode
func (c *Client) SetMirrorMode(enabled bool) error {
	mode := 0
	if enabled {
		mode = 1
	}

	payload := map[string]interface{}{
		"Command": "Device/SetMirrorMode",
		"Mode":    mode,
	}
	_, err := c.sendCommand(payload)
	return err
}

// SetTime24Flag sets the 24-hour time display flag
func (c *Client) SetTime24Flag(enabled bool) error {
	flag := 0
	if enabled {
		flag = 1
	}

	payload := map[string]interface{}{
		"Command": "Device/SetTime24Flag",
		"Mode":    flag,
	}
	_, err := c.sendCommand(payload)
	return err
}

// SetHighLightMode sets the high light mode
// mode: 0 = off, 1 = on
func (c *Client) SetHighLightMode(mode int) error {
	payload := map[string]interface{}{
		"Command": "Device/SetHighLightMode",
		"Mode":    mode,
	}
	_, err := c.sendCommand(payload)
	return err
}

// SetWhiteBalance sets the white balance
// r, g, b: 0-100
func (c *Client) SetWhiteBalance(r, g, b int) error {
	if r < 0 || r > 100 || g < 0 || g > 100 || b < 0 || b > 100 {
		return fmt.Errorf("invalid RGB values (must be 0-100): r=%d, g=%d, b=%d", r, g, b)
	}

	payload := map[string]interface{}{
		"Command": "Device/SetWhiteBalance",
		"RValue":  r,
		"GValue":  g,
		"BValue":  b,
	}
	_, err := c.sendCommand(payload)
	return err
}

// PlayTFGif plays a GIF from TF card
// fileName: name of the GIF file on TF card
func (c *Client) PlayTFGif(fileName string) error {
	payload := map[string]interface{}{
		"Command":  "Device/PlayTFGif",
		"FileType": 2,
		"FileName": fileName,
	}
	_, err := c.sendCommand(payload)
	return err
}

// SaveTFGif saves current display to TF card as GIF
func (c *Client) SaveTFGif() error {
	payload := map[string]string{
		"Command": "Device/SaveTFGif",
	}
	_, err := c.sendCommand(payload)
	return err
}

// PlayBuzzer plays the device buzzer
func (c *Client) PlayBuzzer(params BuzzerParams) error {
	payload := map[string]interface{}{
		"Command":           "Device/PlayBuzzer",
		"ActiveTimeInCycle": params.ActiveTimeInCycle,
		"OffTimeInCycle":    params.OffTimeInCycle,
		"PlayTotalTime":     params.PlayTotalTime,
	}
	_, err := c.sendCommand(payload)
	return err
}

// SetClockFace sets the clock face by ID
func (c *Client) SetClockFace(clockID int) error {
	payload := map[string]interface{}{
		"Command": "Channel/SetClockSelectId",
		"ClockId": clockID,
	}
	_, err := c.sendCommand(payload)
	return err
}

// SetVisualizer sets the audio visualizer equalizer position
func (c *Client) SetVisualizer(equalizerPosition int) error {
	payload := map[string]interface{}{
		"Command":    "Channel/SetEqPosition",
		"EqPosition": equalizerPosition,
	}
	_, err := c.sendCommand(payload)
	return err
}
