package types

type Repository struct {
	FullName        string `json:"full_name,omitempty"`
	HTMLURL         string `json:"html_url,omitempty"`
	ForksCount      int    `json:"forks_count,omitempty"`
	StargazersCount int    `json:"stargazers_count,omitempty"`
	Visibility      string `json:"visibility,omitempty"`
}
