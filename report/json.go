package report

import (
	"time"

	"go.rikki.moe/librespeedtest-cli/defs"
)

// JSONReport represents the output data fields in a JSON file
type JSONReport struct {
	Timestamp     time.Time `json:"timestamp"`
	Server        Server    `json:"server"`
	Client        Client    `json:"client"`
	BytesSent     uint64    `json:"bytes_sent"`
	BytesReceived uint64    `json:"bytes_received"`
	Ping          float64   `json:"ping"`
	Jitter        float64   `json:"jitter"`
	Upload        float64   `json:"upload"`
	Download      float64   `json:"download"`
	Share         string    `json:"share"`
}

// Server represents the speed test server's information
type Server struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Client represents the speed test client's information
type Client struct {
	defs.IPInfoResponse
}
