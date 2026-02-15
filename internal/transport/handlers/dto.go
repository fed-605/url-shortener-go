package handlsers

import resp "github.com/fed-605/url-shortener-go/internal/lib/api/response"

// struct of user request
// Alias field is not required, if the user doesn't give an alias
// the app will generate it automatically
type SaveUrlRequest struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

// struct of app response
type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}
