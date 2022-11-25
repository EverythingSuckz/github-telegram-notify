package main

import (
	"encoding/json"
	"github-telegram-notify/types"
	"github-telegram-notify/utils"
	"io/ioutil"
	"os"
	"testing"

	dotenv "github.com/joho/godotenv"
)

func loadEnvs(t *testing.T) (string, string, string) {
	err := dotenv.Load()
	if err != nil {
		t.Fatal("Error loading .env file")
	}
	tg_token := os.Getenv("BOT_TOKEN")
	if tg_token == "" {
		t.Fatal("Bot token not specified in .env file")
	}
	chatID := os.Getenv("CHAT_ID")
	if chatID == "" {
		t.Fatal("Chat ID not specified in .env file")
	}
	topicID := os.Getenv("TOPIC_ID")
	return tg_token, chatID, topicID
}

func parse(t *testing.T, rawData []byte) (string, string, string) {
	var gitEvent *types.Metadata
	err := json.Unmarshal([]byte(rawData), &gitEvent)
	if err != nil {
		t.Fatal(err)
	}

	text, markupText, markupUrl, err := utils.CreateContents(gitEvent)
	if err != nil {
		t.Fatal(err)
	}
	return text, markupText, markupUrl
}

func TestCommitMessage(t *testing.T) {
	token, chatID, topicID := loadEnvs(t)
	data, err := ioutil.ReadFile("events/commit.json")
	if err != nil {
		t.Fatal(err)
	}
	text, markupText, markupUrl := parse(t, data)
	error := utils.SendMessage(token, chatID, text, markupText, markupUrl, topicID)
	if error.Description != "" {
		t.Fatal(error.String())
	}
}
