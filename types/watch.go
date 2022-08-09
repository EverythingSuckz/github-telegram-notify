package types

// Specifically for star event which requires watch event to be triggered.
type WatchEvent struct {
	Action string     `json:"action,omitempty"`
	Repo   Repository `json:"repository,omitempty"`
	Sender User       `json:"sender,omitempty"`
}
