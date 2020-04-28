package sentiment

// URL of the sentiment-analysis API
const URL = "https://api.deepai.org/api/sentiment-analysis"

// A Classification represents a possible sentimental vibe
// the api will analyze text as
type Classification string

// Classifications returned from the API
const (
	VeryNegative Classification = "Verynegative"
	Negative     Classification = "Negative"
	Neutral      Classification = "Neutral"
	Positive     Classification = "Positive"
	Verypositive Classification = "Verypositive"
)

// Response returns in output containing a slice of classifications describing the
// sentiment of provided text
type Response struct {
	Output []Classification `json:"output"`
	ID     string           `json:"id"`
}
