package recognition

// URL of the face recognition api
const URL = "https://api.deepai.org/api/facial-recognition"

// Response is returned by the deepai face recognition api
type Response struct {
	Output Output `json:"output"`
	ID     string `json:"id"`
}

// The Output of the api containing any detected faces
type Output struct {
	Faces []Face `json:"faces"`
}

// Face contains fields describing the confidence of the
// facial recognition match
type Face struct {
	Confidence  string  `json:"confidence"`
	BoundingBox []int64 `json:"bounding_box"`
	Name        string  `json:"name"`
}
