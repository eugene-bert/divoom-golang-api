package main

import (
	"fmt"
	"log"
	"time"

	divoom "github.com/eugene-bert/divoom-golang-api"
)

func main() {
	// Your PIXOO64's IP address
	deviceIP := "192.168.1.180"

	// Create a new client
	client := divoom.NewClient(deviceIP)

	// IMPORTANT: Switch to Custom channel (3) for custom displays
	fmt.Println("Switching to Custom channel...")
	if err := client.SetChannelIndex(3); err != nil {
		log.Printf("Error switching channel: %v\n", err)
	}
	time.Sleep(500 * time.Millisecond)

	// Example 1: Set brightness
	fmt.Println("Setting brightness to 80...")
	if err := client.SetBrightness(80); err != nil {
		log.Printf("Error setting brightness: %v\n", err)
	}

	// Example 2: Display text (EASY WAY - recommended!)
	fmt.Println("Displaying text (using DisplayText helper)...")
	if err := client.DisplayText("Hello PIXOO64!", "#00FF00"); err != nil {
		log.Printf("Error displaying text: %v\n", err)
	}

	time.Sleep(3 * time.Second)

	// Example 2b: Display text (MANUAL WAY - shows the details)
	fmt.Println("Displaying text (manual way)...")
	// First send a blank screen (text overlays on animations/GIFs)
	if err := client.SendBlankScreen(); err != nil {
		log.Printf("Error sending blank screen: %v\n", err)
	}
	time.Sleep(200 * time.Millisecond)

	// Then send the text
	err := client.SendText(divoom.TextParams{
		TextID:     1,
		X:          0,
		Y:          24,
		Direction:  0,
		Font:       4,
		TextWidth:  64,
		Speed:      10,
		TextString: "Hello PIXOO64!",
		Color:      "#00FF00",
		Align:      2, // center
	})
	if err != nil {
		log.Printf("Error sending text: %v\n", err)
	}

	time.Sleep(3 * time.Second)

	// Example 3: Get all configuration
	fmt.Println("Getting device configuration...")
	config, err := client.GetAllConf()
	if err != nil {
		log.Printf("Error getting config: %v\n", err)
	} else {
		fmt.Printf("Current brightness: %d\n", config.Brightness)
		fmt.Printf("Screen rotation: %d\n", config.GyrateAngle)
		fmt.Printf("24-hour mode: %d\n", config.Time24Flag)
	}

	// Example 4: Set screen rotation
	fmt.Println("Setting screen rotation to 0 degrees...")
	if err := client.SetScreenRotationAngle(0); err != nil {
		log.Printf("Error setting rotation: %v\n", err)
	}

	// Example 5: Get device time
	fmt.Println("Getting device time...")
	deviceTime, err := client.GetDeviceTime()
	if err != nil {
		log.Printf("Error getting time: %v\n", err)
	} else {
		fmt.Printf("Device time: %s\n", deviceTime.LocalTime)
	}

	// Example 6: Set a timer
	fmt.Println("Setting a 1 minute timer...")
	err = client.SetTimer(divoom.TimerParams{
		Minute: 1,
		Second: 0,
		Status: 1, // start
	})
	if err != nil {
		log.Printf("Error setting timer: %v\n", err)
	}

	// Example 7: Set scoreboard
	fmt.Println("Setting scoreboard...")
	err = client.SetScoreBoard(divoom.ScoreBoardParams{
		BlueScore: 5,
		RedScore:  3,
	})
	if err != nil {
		log.Printf("Error setting scoreboard: %v\n", err)
	}

	time.Sleep(2 * time.Second)

	// Example 8: Set location for weather
	fmt.Println("Setting location...")
	if err := client.SetLogAndLat("30.29", "20.58"); err != nil {
		log.Printf("Error setting location: %v\n", err)
	}

	// Example 9: Get weather info
	fmt.Println("Getting weather info...")
	weather, err := client.GetWeatherInfo()
	if err != nil {
		log.Printf("Error getting weather: %v\n", err)
	} else {
		fmt.Printf("Weather: %s\n", weather.Weather)
		fmt.Printf("Temperature: %.1f°C\n", weather.CurTemp)
		fmt.Printf("Humidity: %d%%\n", weather.Humidity)
	}

	// Example 10: Play buzzer
	fmt.Println("Playing buzzer...")
	err = client.PlayBuzzer(divoom.BuzzerParams{
		ActiveTimeInCycle: 500,  // 500ms on
		OffTimeInCycle:    500,  // 500ms off
		PlayTotalTime:     2000, // 2 seconds total
	})
	if err != nil {
		log.Printf("Error playing buzzer: %v\n", err)
	}

	time.Sleep(3 * time.Second)

	// Example 11: Clear text
	fmt.Println("Clearing text...")
	if err := client.ClearText(); err != nil {
		log.Printf("Error clearing text: %v\n", err)
	}

	fmt.Println("All examples completed!")
}
