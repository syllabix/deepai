package config

import (
	"net/http"
	"time"
)

// An Option is used to configure a setting for a deepai client
type Option func(*Settings)

// Settings are used to configure a deepai client
type Settings struct {
	Timeout      time.Duration
	RoundTripper http.RoundTripper
}
