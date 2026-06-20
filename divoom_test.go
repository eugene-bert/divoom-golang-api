package divoom

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient("192.168.1.100")

	if c.deviceIP != "192.168.1.100" {
		t.Errorf("deviceIP = %q, want %q", c.deviceIP, "192.168.1.100")
	}
	if c.baseURL != "http://192.168.1.100:80/post" {
		t.Errorf("baseURL = %q, want %q", c.baseURL, "http://192.168.1.100:80/post")
	}
	if c.httpClient == nil {
		t.Error("httpClient is nil")
	}
}

func TestGetDeviceIP(t *testing.T) {
	c := NewClient("10.0.0.5")
	if got := c.GetDeviceIP(); got != "10.0.0.5" {
		t.Errorf("GetDeviceIP() = %q, want %q", got, "10.0.0.5")
	}
}

func TestParseHexColor(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"#ff0000", "#FF0000"},
		{"ff0000", "#FF0000"},
		{"#FF0000", "#FF0000"},
		{" ff0000 ", "#FF0000"},
		{"#abc", "#ABC"},
	}

	for _, tt := range tests {
		got := ParseHexColor(tt.input)
		if got != tt.want {
			t.Errorf("ParseHexColor(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
