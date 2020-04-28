package deepai

import (
	"context"
	"encoding/json"
	"io"

	"github.com/syllabix/deepai/recognition"
)

// RecognizeFaces will use the deepai api to recognize faces in the provided image
// The method will return a non empty recognition response on success, or a non nil error on failure
func (c *Client) RecognizeFaces(ctx context.Context, file io.Reader) (recognition.Response, error) {
	var response recognition.Response

	mpw, img := imageForm(file)
	result, err := c.request(ctx, recognition.URL, mpw.FormDataContentType(), img)
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
