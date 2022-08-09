package types

type PullRequest struct {
	Number  int    `json:"number,omitempty"`
	State   string `json:"state,omitempty"`
	Locked  bool   `json:"locked,omitempty"`
	Title   string `json:"title,omitempty"`
	User    User   `json:"user,omitempty"`
	HTMLURL string `json:"html_url,omitempty"`
}

type PullRequestEvent struct {
	Action      string       `json:"action,omitempty"`
	Number      int          `json:"number,omitempty"`
	PullRequest *PullRequest `json:"pull_request,omitempty"`
	Repo        Repository   `json:"repository,omitempty"`
	Sender      User         `json:"sender,omitempty"`
}

type PullRequestReviewCommentEvent struct {
	Action      string              `json:"action,omitempty"`
	PullRequest *PullRequest        `json:"pull_request,omitempty"`
	Comment     *PullRequestComment `json:"comment,omitempty"`
	Repo        Repository          `json:"repository,omitempty"`
	Sender      User                `json:"sender,omitempty"`
}

type PullRequestComment struct {
	User           User   `json:"user,omitempty"`
	HTMLURL        string `json:"html_url,omitempty"`
	PullRequestURL string `json:"pull_request_url,omitempty"`
}
