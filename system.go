package divoom

import "fmt"

// SetLogAndLat sets the longitude and latitude for weather information
func (c *Client) SetLogAndLat(longitude, latitude string) error {
	payload := map[string]interface{}{
		"Command":   "Sys/LogAndLat",
		"Longitude": longitude,
		"Latitude":  latitude,
	}
	_, err := c.sendCommand(payload)
	return err
}

// SetTimeZone sets the device time zone
// timeZone: time zone string, e.g., "America/New_York"
func (c *Client) SetTimeZone(timeZone string) error {
	payload := map[string]interface{}{
		"Command":  "Sys/TimeZone",
		"TimeZone": timeZone,
	}
	_, err := c.sendCommand(payload)
	return err
}

// GetWeatherInfo gets the current weather information
func (c *Client) GetWeatherInfo() (*WeatherInfo, error) {
	payload := map[string]string{
		"Command": "Device/GetWeatherInfo",
	}

	var result WeatherInfo
	err := c.sendCommandWithResponse(payload, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrorCode != 0 {
		return nil, fmt.Errorf("device returned error code: %d", result.ErrorCode)
	}

	return &result, nil
}
