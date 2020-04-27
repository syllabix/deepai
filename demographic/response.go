package demographic

const (
	// URL is the url of the demographic recognition api
	URL = "https://api.deepai.org/api/demographic-recognition"
)

// Response contains the output from a deepai demographic recognition analysis
type Response struct {
	Output Output `json:"output"`
	ID     string `json:"id"`
}

// Output contains a slice of demographically recognized faces
type Output struct {
	Faces []Face `json:"faces"`
}

// Face contains demographic information about the detected fahce as well as its
// location in the image
type Face struct {
	AgeRange                     []int64 `json:"age_range"`
	CulturalAppearanceConfidence float64 `json:"cultural_appearance_confidence"`
	Gender                       string  `json:"gender"`
	AgeRangeConfidence           float64 `json:"age_range_confidence"`
	BoundingBox                  []int64 `json:"bounding_box"`
	GenderConfidence             float64 `json:"gender_confidence"`
	CulturalAppearance           string  `json:"cultural_appearance"`
}
