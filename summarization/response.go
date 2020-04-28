package summarization

// URL of the summarization api
const URL = "https://api.deepai.org/api/summarization"

// Response contains output containing a summarized version of the provided
// text
type Response struct {
	Output string `json:"output"`
	ID     string `json:"id"`
}
