package divoom

// StandardResponse represents a standard API response
type StandardResponse struct {
	ErrorCode int `json:"error_code"`
}

// CommandRequest represents a basic command request
type CommandRequest struct {
	Command string `json:"Command"`
}

// TextParams represents parameters for sending text to display
type TextParams struct {
	TextID     int    `json:"TextId"`
	X          int    `json:"x"`
	Y          int    `json:"y"`
	Direction  int    `json:"dir"`  // 0: left, 1: right
	Font       int    `json:"font"` // Font ID
	TextWidth  int    `json:"TextWidth"`
	Speed      int    `json:"speed"` // Scroll speed in ms
	TextString string `json:"TextString"`
	Color      string `json:"color"` // Hex color, e.g., "#FFFF00"
	Align      int    `json:"align"` // 1: left, 2: middle, 3: right
}

// GifParams represents parameters for sending GIF data
type GifParams struct {
	PicNum    int    `json:"PicNum"`    // Number of frames
	PicWidth  int    `json:"PicWidth"`  // Width (should be 64 for PIXOO64)
	PicOffset int    `json:"PicOffset"` // Frame offset
	PicID     int    `json:"PicID"`     // Picture ID
	PicSpeed  int    `json:"PicSpeed"`  // Animation speed in ms
	PicData   string `json:"PicData"`   // Base64 encoded image data
}

// TimerParams represents parameters for setting a timer
type TimerParams struct {
	Minute int `json:"Minute"`
	Second int `json:"Second"`
	Status int `json:"Status"` // 0: stop, 1: start
}

// StopWatchParams represents parameters for stopwatch
type StopWatchParams struct {
	Status int `json:"Status"` // 0: stop, 1: start, 2: reset
}

// ScoreBoardParams represents parameters for scoreboard
type ScoreBoardParams struct {
	BlueScore int `json:"BlueScore"`
	RedScore  int `json:"RedScore"`
}

// AllConfigResponse represents the response from GetAllConf
type AllConfigResponse struct {
	ErrorCode           int `json:"error_code"`
	Brightness          int `json:"Brightness"`
	RotationFlag        int `json:"RotationFlag"`
	ClockTime           int `json:"ClockTime"`
	GalleryTime         int `json:"GalleryTime"`
	SingleGalleyTime    int `json:"SingleGalleyTime"`
	PowerOnChannelId    int `json:"PowerOnChannelId"`
	GalleryShowTimeFlag int `json:"GalleryShowTimeFlag"`
	CurClockId          int `json:"CurClockId"`
	Time24Flag          int `json:"Time24Flag"`
	TemperatureMode     int `json:"TemperatureMode"`
	GyrateAngle         int `json:"GyrateAngle"`
	MirrorFlag          int `json:"MirrorFlag"`
	LightSwitch         int `json:"LightSwitch"`
}

// WeatherInfo represents weather information
type WeatherInfo struct {
	ErrorCode  int     `json:"error_code"`
	Weather    string  `json:"Weather"`
	CurTemp    float64 `json:"CurTemp"`
	MinTemp    float64 `json:"MinTemp"`
	MaxTemp    float64 `json:"MaxTemp"`
	Pressure   int     `json:"Pressure"`
	Humidity   int     `json:"Humidity"`
	Visibility int     `json:"Visibility"`
	WindSpeed  float64 `json:"WindSpeed"`
}

// DeviceTime represents device time information
type DeviceTime struct {
	ErrorCode int    `json:"error_code"`
	UTCTime   int64  `json:"UTCTime"`
	LocalTime string `json:"LocalTime"`
}

// BuzzerParams represents parameters for playing buzzer
type BuzzerParams struct {
	ActiveTimeInCycle int `json:"ActiveTimeInCycle"` // Active time in ms
	OffTimeInCycle    int `json:"OffTimeInCycle"`    // Off time in ms
	PlayTotalTime     int `json:"PlayTotalTime"`     // Total play time in ms
}

// LocationParams represents location for weather
type LocationParams struct {
	Longitude string `json:"Longitude"`
	Latitude  string `json:"Latitude"`
}
