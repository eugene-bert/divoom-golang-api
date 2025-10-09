package divoom

// SetTimer sets a countdown timer
func (c *Client) SetTimer(params TimerParams) error {
	payload := map[string]interface{}{
		"Command": "Tools/SetTimer",
		"Minute":  params.Minute,
		"Second":  params.Second,
		"Status":  params.Status,
	}
	_, err := c.sendCommand(payload)
	return err
}

// SetStopWatch controls the stopwatch
// status: 0 = stop, 1 = start, 2 = reset
func (c *Client) SetStopWatch(params StopWatchParams) error {
	payload := map[string]interface{}{
		"Command": "Tools/SetStopWatch",
		"Status":  params.Status,
	}
	_, err := c.sendCommand(payload)
	return err
}

// SetScoreBoard sets the scoreboard display
func (c *Client) SetScoreBoard(params ScoreBoardParams) error {
	payload := map[string]interface{}{
		"Command":   "Tools/SetScoreBoard",
		"BlueScore": params.BlueScore,
		"RedScore":  params.RedScore,
	}
	_, err := c.sendCommand(payload)
	return err
}

// SetNoiseStatus controls the noise detection status
// status: 0 = off, 1 = on
func (c *Client) SetNoiseStatus(status int) error {
	payload := map[string]interface{}{
		"Command":     "Tools/SetNoiseStatus",
		"NoiseStatus": status,
	}
	_, err := c.sendCommand(payload)
	return err
}
