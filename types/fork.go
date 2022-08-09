package types

type ForkEvent struct {
	Forkee Repository `json:"forkee,omitempty"`
	Repo   Repository `json:"repository,omitempty"`
	Sender User       `json:"sender,omitempty"`
}
