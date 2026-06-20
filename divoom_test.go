package divoom

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

func newTestServer(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	c := NewClient("127.0.0.1")
	c.baseURL = srv.URL
	return c
}

func TestSendCommandRoundTrip(t *testing.T) {
	var gotCommand string
	c := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		var payload map[string]interface{}
		json.NewDecoder(r.Body).Decode(&payload)
		gotCommand, _ = payload["Command"].(string)
		json.NewEncoder(w).Encode(map[string]int{"error_code": 0})
	})

	err := c.SetBrightness(50)
	if err != nil {
		t.Fatalf("SetBrightness: %v", err)
	}
	if gotCommand != "Channel/SetBrightness" {
		t.Errorf("command = %q, want Channel/SetBrightness", gotCommand)
	}
}

func TestSendCommandDeviceError(t *testing.T) {
	c := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]int{"error_code": 42})
	})

	err := c.SetBrightness(50)
	if err == nil {
		t.Fatal("expected error for non-zero error_code")
	}
}

func TestSendCommandHTTPError(t *testing.T) {
	c := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	err := c.SetBrightness(50)
	if err == nil {
		t.Fatal("expected error for HTTP 500")
	}
}

func TestSendCommandBadJSON(t *testing.T) {
	c := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not json"))
	})

	err := c.SetBrightness(50)
	if err == nil {
		t.Fatal("expected error for bad JSON response")
	}
}

func TestSetTimeoutConcurrent(t *testing.T) {
	c := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]int{"error_code": 0})
	})

	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := 0; i < 100; i++ {
			c.SetTimeout(time.Duration(i+1) * time.Second)
		}
	}()
	for i := 0; i < 100; i++ {
		c.SetBrightness(50)
	}
	<-done
}

func TestSendCommandWithResponse(t *testing.T) {
	c := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error_code": 0,
			"Brightness": 75,
			"CurClockId": 100,
			"Time24Flag": 1,
		})
	})

	conf, err := c.GetAllConf()
	if err != nil {
		t.Fatalf("GetAllConf: %v", err)
	}
	if conf.Brightness != 75 {
		t.Errorf("Brightness = %d, want 75", conf.Brightness)
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
