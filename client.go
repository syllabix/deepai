package deepai

import (
	"context"
	"deepai/config"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	apiKey = "api-key"
)

// common client errors
var (
	ErrUnauthorized = errors.New("the provided api token is either not valid or not authorized to make this request")
	ErrBadRequest   = errors.New("an error occurred while processing given inputs from the request")
	ErrUnknown      = errors.New("the server responded with an unknown error")
)

// Client implements methods that can be used to interface with the deepai API
// The client should be reused instead of created as needed and safe concurrent
// use by multiple go routines
type Client struct {
	token string
	http  *http.Client
}

func (c *Client) request(ctx context.Context, url string, contentType string, body io.Reader) (io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set(apiKey, c.token)
	req = req.WithContext(ctx)
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		return res.Body, nil
	}

	defer res.Body.Close()
	var msg strings.Builder
	_, err = io.Copy(&msg, res.Body)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusBadRequest:
		return nil, fmt.Errorf("%w: %s", ErrBadRequest, msg.String())

	case http.StatusUnauthorized:
		return nil, fmt.Errorf("%w: %s", ErrUnauthorized, msg.String())

	default:
		return nil, ErrUnknown
	}
}

// NewClient takes the provided API Key and a variadic list of options
// and returns a ready to use deepai Client.
// For more information on how to obtain a token
// please visit https://https://deepai.org/
func NewClient(apiKey string, options ...config.Option) *Client {

	// set up defaults
	settings := &config.Settings{
		Timeout:      time.Second * 15,
		RoundTripper: http.DefaultTransport,
	}

	for _, opt := range options {
		opt(settings)
	}

	return &Client{
		token: apiKey,
		http: &http.Client{
			Timeout:   settings.Timeout,
			Transport: settings.RoundTripper,
		},
	}
}
