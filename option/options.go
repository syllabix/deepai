package option

import (
	"net/http"
	"time"

	"github.com/syllabix/deepai/config"
)

// Timeout is a config option for a deepai client that will set
// the request timeout duration
func Timeout(t time.Duration) config.Option {
	return func(s *config.Settings) {
		s.Timeout = t
	}
}

// Transport is a config option for a deepai client that will set
// the underlying http RoundTripper the client will use
func Transport(t http.RoundTripper) config.Option {
	return func(s *config.Settings) {
		s.RoundTripper = t
	}
}
