package main

import (
	"encoding/json"
	"github-telegram-notify/types"
	"github-telegram-notify/utils"
	"os"
)

func main() {
	tg_token := os.Getenv("INPUT_BOT_TOKEN")
	chatID := os.Getenv("INPUT_CHAT_ID")
	topicID := os.Getenv("INPUT_TOPIC_ID")
	gitEventRaw := os.Getenv("INPUT_GIT_EVENT")
	print(gitEventRaw)
	var gitEvent *types.Metadata
	err := json.Unmarshal([]byte(gitEventRaw), &gitEvent)
	if err != nil {
		panic(err)
	}
	text, markupText, markupUrl, err := utils.CreateContents(gitEvent)
	if err != nil {
		panic(err)
	}
	error := utils.SendMessage(tg_token, chatID, text, markupText, markupUrl, topicID)
	if error.Description != "" {
		panic(error.String())
	}

}
