package deepai

import (
	"context"
	"encoding/json"
	"io"

	"github.com/syllabix/deepai/moderation"
)

// Moderate the provided image content using the deepai content moderation API
// The method will return a moderation response on success, or a non nil error on failure
func (c *Client) Moderate(ctx context.Context, file io.Reader) (moderation.Response, error) {
	var response moderation.Response

	mpw, img := imageForm(file)
	result, err := c.request(ctx, moderation.URL, mpw.FormDataContentType(), img)
	if err != nil {
		return response, err
	}
	defer result.Close()

	err = json.NewDecoder(result).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}
