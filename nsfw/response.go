package nsfw

const (
	// URL is the url of the content moderation api
	URL = "https://api.deepai.org/api/nsfw-detector"
)

// The Detectable nsfw component of an image
type Detectable string

// Available NSFW Features
const (
	ExposedFemaleGenitalia Detectable = "Female Genitalia - Exposed"
	CoveredFemaleGenitalia Detectable = "Female Genitalia - Covered"
	ExposedFemaleBreast    Detectable = "Female Breast - Exposed"
	CoveredFemaleBreast    Detectable = "Female Breast - Covered"
	ExposedMaleGenitalia   Detectable = "Male Genitalia - Exposed"
	CoveredMaleGenitalia   Detectable = "Male Genitalia - Covered"
	ExposedMaleBreast      Detectable = "Male Breast - Exposed"
	CoveredMaleBreast      Detectable = "Male Breast - Covered"
	ExposedButtocks        Detectable = "Buttocks - Exposed"
)

// Response contains the output and id of a content moderation response
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
	Confidence  string     `json:"confidence"`
	BoundingBox []int64    `json:"bounding_box"`
	Name        Detectable `json:"name"`
}
