package discordtexthook

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func sendAPIRequest(message, messageID, todo string, ds discord, embed *Embed) (*http.Response, error) {
	var method, apiURL string

	switch todo {
	case "send", "send_embed":
		method = "POST"
		apiURL = "https://discord.com/api/v9/webhooks/" + ds.WebhookID + "/" + ds.WebhookToken + "?wait=true"
	case "edit", "edit_embed":
		method = "PATCH"
		apiURL = "https://discord.com/api/v9/webhooks/" + ds.WebhookID + "/" + ds.WebhookToken + "/messages/" + messageID
	case "delete":
		method = "DELETE"
		apiURL = "https://discord.com/api/v9/webhooks/" + ds.WebhookID + "/" + ds.WebhookToken + "/messages/" + messageID
	}

	var payload interface{}
	if embed != nil {
		payload = WebhookPayload{
			Embeds: []Embed{*embed},
		}
	} else {
		payload = WebhookPayload{
			Content: message,
		}
	}

	byteValue, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	pload := bytes.NewBuffer(byteValue)

	return executeRequest(method, apiURL, pload)git 
}

func executeRequest(method, apiURL string, payload io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, apiURL, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}
