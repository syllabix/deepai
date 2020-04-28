package deepai

import (
	"context"
	"encoding/json"
	"io"

	"github.com/syllabix/deepai/nsfw"
)

// DetectNudity will use the deepai api to detect nudity in the provided file
// The method will return a nsfw moderation response on success, or a non nil error on failure
func (c *Client) DetectNudity(ctx context.Context, file io.Reader) (nsfw.Response, error) {
	var response nsfw.Response

	mpw, img := imageForm(file)
	result, err := c.request(ctx, nsfw.URL, mpw.FormDataContentType(), img)
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
