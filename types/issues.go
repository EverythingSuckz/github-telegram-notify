package types

type Issue struct {
	Number            int        `json:"number,omitempty"`
	State             string     `json:"state,omitempty"`
	Locked            bool       `json:"locked,omitempty"`
	Title             string     `json:"title,omitempty"`
	Body              string     `json:"body,omitempty"`
	AuthorAssociation string     `json:"author_association,omitempty"`
	User              User       `json:"user,omitempty"`
	HTMLURL           string     `json:"html_url,omitempty"`
	Repository        Repository `json:"repository,omitempty"`
}

type IssueCommentEvent struct {
	// Possible values are: "created", "edited", "deleted".
	Action  string        `json:"action,omitempty"`
	Issue   *Issue        `json:"issue,omitempty"`
	Comment *IssueComment `json:"comment,omitempty"`
	Repo    Repository    `json:"repository,omitempty"`
	Sender  User          `json:"sender,omitempty"`
}

type IssueComment struct {
	User     User   `json:"user,omitempty"`
	HTMLURL  string `json:"html_url,omitempty"`
	IssueURL string `json:"issue_url,omitempty"`
}

type IssuesEvent struct {
	Action string     `json:"action,omitempty"`
	Issue  *Issue     `json:"issue,omitempty"`
	Repo   Repository `json:"repository,omitempty"`
	Sender User       `json:"sender,omitempty"`
}
