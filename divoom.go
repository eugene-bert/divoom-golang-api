// Package divoom provides a Go client for the Divoom PIXOO64 LED display.
// It supports text display, canvas drawing, image/GIF rendering, device
// discovery, and full device control with zero external dependencies.
package divoom

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// Client represents a Divoom PIXOO64 API client
type Client struct {
	deviceIP   string
	httpClient *http.Client
	baseURL    string
	mu         sync.Mutex
}

// NewClient creates a new Divoom API client
// deviceIP should be the local IP address of your PIXOO64 device
func NewClient(deviceIP string) *Client {
	return &Client{
		deviceIP: deviceIP,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: fmt.Sprintf("http://%s:80/post", deviceIP),
	}
}

// SetTimeout sets the HTTP client timeout
func (c *Client) SetTimeout(timeout time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.httpClient = &http.Client{Timeout: timeout}
}

func (c *Client) sendCommand(payload interface{}) (*StandardResponse, error) {
	return c.sendCommandCtx(context.Background(), payload)
}

func (c *Client) sendCommandCtx(ctx context.Context, payload interface{}) (*StandardResponse, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected HTTP status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var result StandardResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if result.ErrorCode != 0 {
		return &result, fmt.Errorf("device returned error code: %d", result.ErrorCode)
	}

	return &result, nil
}

func (c *Client) sendCommandWithResponse(payload interface{}, result interface{}) error {
	return c.sendCommandWithResponseCtx(context.Background(), payload, result)
}

func (c *Client) sendCommandWithResponseCtx(ctx context.Context, payload interface{}, result interface{}) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.doRequest(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected HTTP status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	var stdResp StandardResponse
	if err := json.Unmarshal(body, &stdResp); err == nil && stdResp.ErrorCode != 0 {
		return fmt.Errorf("device returned error code: %d", stdResp.ErrorCode)
	}

	return nil
}

func (c *Client) doRequest(req *http.Request) (*http.Response, error) {
	c.mu.Lock()
	client := c.httpClient
	c.mu.Unlock()
	return client.Do(req)
}

// GetDeviceIP returns the device IP address
func (c *Client) GetDeviceIP() string {
	return c.deviceIP
}
