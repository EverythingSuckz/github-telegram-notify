package types

import (
	"encoding/json"
)

type Metadata struct {
	Sha            string           `json:"sha"`
	RepositoryName string           `json:"repository"`
	RawEvent       *json.RawMessage `json:"event"`
	Ref_name       string           `json:"ref_name"`
	ServerUrl      string           `json:"server_url"`
	EventName      string           `json:"event_name"`
}

func (e *Metadata) ParseEvent() (event_type interface{}, err error) {
	switch e.EventName {
	case "fork":
		event_type = &ForkEvent{}
	case "issue_comment":
		event_type = &IssueCommentEvent{}
	case "issues":
		event_type = &IssuesEvent{}
	case "pull_request_target":
		event_type = &PullRequestEvent{}
	case "pull_request_review_comment":
		event_type = &PullRequestReviewCommentEvent{}
	case "push":
		event_type = &PushEvent{}
	case "release":
		event_type = &ReleaseEvent{}
	case "watch":
		event_type = &WatchEvent{}
	}
	err = json.Unmarshal(*e.RawEvent, &event_type)
	return event_type, err
}
