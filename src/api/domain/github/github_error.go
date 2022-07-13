package github

type GitHErrorResponse struct {
	StatusCode       int         `json:"status_code"` //Not in body response
	Message          string      `json:"message"`
	Errors           []GitHError `json:"errors"`
	DocumentationURL string      `json:"documentation_url"`
}

type GitHError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
