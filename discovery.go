package divoom

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const discoveryURL = "https://app.divoom-gz.com/Device/ReturnSameLANDevice"

// DiscoveredDevice represents a Divoom device found on the local network.
type DiscoveredDevice struct {
	DeviceName string `json:"DeviceName"`
	DeviceID   int    `json:"DeviceId"`
	DeviceIP   string `json:"DevicePrivateIP"`
	DeviceMAC  string `json:"DeviceMac"`
	Hardware   int    `json:"Hardware"`
}

// DiscoverDevices finds Divoom devices on the local network via the Divoom cloud API.
// Requires internet access — the device list is returned by Divoom's servers
// based on your public IP and local network.
func DiscoverDevices() ([]DiscoveredDevice, error) {
	return DiscoverDevicesContext(context.Background())
}

// DiscoverDevicesContext is like DiscoverDevices but accepts a context for cancellation.
func DiscoverDevicesContext(ctx context.Context) ([]DiscoveredDevice, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", discoveryURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create discovery request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{Timeout: 10 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("discovery request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("discovery returned HTTP %d", resp.StatusCode)
	}

	var result struct {
		ReturnCode int                `json:"ReturnCode"`
		DeviceList []DiscoveredDevice `json:"DeviceList"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode discovery response: %w", err)
	}

	if result.ReturnCode != 0 {
		return nil, fmt.Errorf("discovery returned error code: %d", result.ReturnCode)
	}

	return result.DeviceList, nil
}
