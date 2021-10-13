package discordtexthook

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func sendAPIRequest(message, messageID, todo string, ds discord) (*http.Response, error) {
	var method, apiURL string

	switch todo {
	case "send":
		method = "POST"
		apiURL = "https://discord.com/api/v9/webhooks/" + ds.WebhookID + "/" + ds.WebhookToken + "?wait=true"
	case "edit":
		method = "PATCH"
		apiURL = "https://discord.com/api/v9/webhooks/" + ds.WebhookID + "/" + ds.WebhookToken + "/messages/" + messageID
	case "delete":
		method = "DELETE"
		apiURL = "https://discord.com/api/v9/webhooks/" + ds.WebhookID + "/" + ds.WebhookToken + "/messages/" + messageID
	}

	payload := map[string]string{"content": message}
	byteValue, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	pload := bytes.NewBuffer(byteValue)

	return executeRequest(method, apiURL, pload)
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
