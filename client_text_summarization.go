package deepai

import (
	"context"
	"encoding/json"
	"io"

	"github.com/syllabix/deepai/summarization"
)

// SummarizeText will use the deepai api to summarize text in the provided file
// The method will return a non empty summarization response on success, or a non nil error on failure
func (c *Client) SummarizeText(ctx context.Context, file io.Reader) (summarization.Response, error) {
	var response summarization.Response

	mpw, img := textForm(file)
	result, err := c.request(ctx, summarization.URL, mpw.FormDataContentType(), img)
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
