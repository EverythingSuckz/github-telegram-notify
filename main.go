package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type Metadata struct {
	Sha            string `json:"sha"`
	RepositoryName string `json:"repository"`
	Event          Event  `json:"event"`
	Ref_name       string `json:"ref_name"`
	ServerUrl      string `json:"server_url"`
}

func (m Metadata) repoUrl() string {
	return fmt.Sprint(m.ServerUrl, "/", m.RepositoryName)
}

type Event struct {
	Commits []Commit `json:"commits"`
	Compare string   `json:"compare"`
}

type Commit struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Url     string `json:"url"`
	Ref     string `json:"ref"`
	Author  Author `json:"author"`
}

type Author struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (a Author) Url() string {
	return fmt.Sprint("https://github.com/", a.Username)
}

func main() {
	apiBaseUri, _ := url.Parse("https://api.telegram.org")
	tg_token := os.Getenv("INPUT_BOT_TOKEN")
	chatId := os.Getenv("INPUT_CHAT_ID")
	gitEventRaw := os.Getenv("INPUT_GIT_EVENT")
	print(gitEventRaw)
	var gitEvent Metadata
	err := json.Unmarshal([]byte(gitEventRaw), &gitEvent)
	if err != nil {
		panic(err)
	}
	req_url, _ := url.Parse(fmt.Sprint(apiBaseUri, "/bot", tg_token, "/sendMessage"))
	params := url.Values{}
	text := fmt.Sprintf("<b>🔨 %d New commit to</b> <a href=\"%s\">%s</a>[<code>%s</code>]\n\n", len(gitEvent.Event.Commits), gitEvent.repoUrl(), gitEvent.RepositoryName, gitEvent.Ref_name)
	for _, commit := range gitEvent.Event.Commits {
		text += fmt.Sprintf("• <a href=\"%s\">%s</a> - %s by <a href=\"%s\">%s</a>\n", commit.Url, commit.Id[:7], commit.Message, commit.Author.Url(), commit.Author.Name)
	}
	params.Set("chat_id", chatId)
	params.Set("text", text)
	params.Set("parse_mode", "html")
	params.Set("disable_web_page_preview", "true")
	kyb, err := json.Marshal(map[string][][]map[string]string{
		"inline_keyboard": {
			{{"text": "Open changes", "url": gitEvent.Event.Compare}},
		},
	})
	if err != nil {
		panic(err)
	}
	params.Set("reply_markup", string(kyb))
	req_url.RawQuery = params.Encode()
	http.Get(req_url.String())
}
