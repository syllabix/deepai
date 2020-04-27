package moderation

const (
	// URL is the url of the content moderation api
	URL = "https://api.deepai.org/api/content-moderation"
)

// Response contains the ouput and id of a content moderation response
type Response struct {
	Output Output `json:"output"`
	ID     string `json:"id"`
}

// Output contains the detections and "not safe for work"
// score of a provided image
type Output struct {
	Detections []Detection `json:"detections"`
	NsfwScore  float64     `json:"nsfw_score"`
}

// Detection is the summary of a potentially sensitive
// section of an image
type Detection struct {
	Confidence  string  `json:"confidence"`
	BoundingBox []int64 `json:"bounding_box"`
	Name        string  `json:"name"`
}
