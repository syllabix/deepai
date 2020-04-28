package deepai

import (
	"context"
	"deepai/sentiment"
	"encoding/json"
	"io"
)

// AnalyzeSentiment will use the deepai api to anaylse the sentiment of text in the provided file
// The method will return a non empty sentiment response on success, or a non nil error on failure
func (c *Client) AnalyzeSentiment(ctx context.Context, file io.Reader) (sentiment.Response, error) {
	var response sentiment.Response

	mpw, img := textForm(file)
	result, err := c.request(ctx, sentiment.URL, mpw.FormDataContentType(), img)
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
