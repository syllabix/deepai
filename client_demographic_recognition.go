package deepai

import (
	"context"
	"encoding/json"
	"io"

	"github.com/syllabix/deepai/demographic"
)

// GetFaceDemographics will use the deepai api to detect faces and predict associated demographics in the provided image
// The method will return a non empty demographic response on success, or a non nil error on failure
func (c *Client) GetFaceDemographics(ctx context.Context, file io.Reader) (demographic.Response, error) {
	var response demographic.Response

	mpw, img := imageForm(file)
	result, err := c.request(ctx, demographic.URL, mpw.FormDataContentType(), img)
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
