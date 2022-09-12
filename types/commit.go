package types

type PushEvent struct {
	Ref     string     `json:"ref,omitempty"`
	Commits []Commit   `json:"commits,omitempty"`
	Action  string     `json:"action,omitempty"`
	Repo    Repository `json:"repository,omitempty"`
	Compare string     `json:"compare,omitempty"`
}

type Commit struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Url     string `json:"url"`
	Ref     string `json:"ref"`
	Author  User   `json:"author"`
}
