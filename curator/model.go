package curator

type Recommendation string

const (
	Unknown        Recommendation = "unknown"
	Recommended    Recommendation = "0"
	NotRecommended Recommendation = "1"
	Informative    Recommendation = "2"
)

type Review struct {
	AppId          string         `json:"appId"`
	ReviewContent  string         `json:"content"`
	FullReviewUrl  string         `json:"fullReview"`
	Recommendation Recommendation `json:"recommendation"`
}

type GetReviewsResponse struct {
	Success     int    `json:"success"`
	PageSize    string `json:"pagesize"`
	TotalCount  int    `json:"total_count"`
	Start       string `json:"start"`
	ResultsHtml string `json:"results_html"`
}
