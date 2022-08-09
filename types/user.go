package types

type User struct {
	Name    string `json:"name,omitempty"`
	Login   string `json:"login,omitempty"`
	HTMLURL string `json:"html_url,omitempty"`
}
