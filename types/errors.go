package types

import "fmt"

type Error struct {
	Module      string
	Description string
	Message     string
}

func (e Error) String() string {
	return fmt.Sprintf("%s: %s\n%s", e.Module, e.Description, e.Message)
}

type TelegramError struct {
	OK          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}
