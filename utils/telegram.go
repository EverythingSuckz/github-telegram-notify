package utils

import (
	"encoding/json"
	"fmt"
	"github-telegram-notify/types"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendMessage(token string, chatID string, text string, markupText string, markupUrl string) (error types.Error) {
	apiBaseUri, _ := url.Parse("https://api.telegram.org")
	req_url, _ := url.Parse(fmt.Sprint(apiBaseUri, "/bot", token, "/sendMessage"))
	params := url.Values{}
	params.Set("chat_id", chatID)
	escaped_text := html.EscapeString(text)
	params.Set("text", escaped_text)
	params.Set("parse_mode", "html")
	params.Set("disable_web_page_preview", "true")
	kyb, err := json.Marshal(map[string][][]map[string]string{
		"inline_keyboard": {
			{{"text": markupText, "url": markupUrl}},
		},
	})
	if err != nil {
		return types.Error{
			Module:      "json",
			Description: "Failed to marshal inline keyboard",
			Message:     err.Error(),
		}
	}
	params.Set("reply_markup", string(kyb))
	req_url.RawQuery = params.Encode()
	resp, err := http.Get(req_url.String())
	if err != nil {
		return types.Error{
			Module:      "http",
			Description: "Failed to send message",
			Message:     err.Error(),
		}
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		var tgErr types.TelegramError
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return types.Error{
				Module:      "ioutil",
				Description: "Failed to read response body",
				Message:     err.Error(),
			}
		}
		if err := json.Unmarshal(body, &tgErr); err != nil {
			return types.Error{
				Module:      "json",
				Description: "Failed to unmarshal response body",
				Message:     err.Error(),
			}
		}
		return types.Error{
			Module:      "telegram",
			Description: "Failed to send message",
			Message:     tgErr.Description,
		}
	}
	return
}
