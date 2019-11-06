package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type WebhookMessage struct {
	Channel  string   `json:"channel"`
	UserName string   `json:"username"`
	Text     string   `json:"text"`
	Mrkdwn   bool     `json:"mrkdwn"`
	Blocks   []*Block `json:"blocks"`
}

type Block struct {
	Type string `json:"type"`
	Text Text   `json:"text"`
}

type Text struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji"`
}

func PostWebhookMessage(url string, msg *WebhookMessage) (*http.Response, error) {
	b, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return res, nil
}
