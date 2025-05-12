package discordtexthook

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type discordInterfacer interface {
	SendMessage(message string) (*Message, error)
	SendEmbed(embed Embed) (*Message, error)
	EditMessage(messageID, message string) (*Message, error)
	EditEmbed(messageID string, embed Embed) (*Message, error)
	DeleteMessage(messageID string) error
}

type discord struct {
	WebhookID    string
	WebhookToken string
}

func (ds discord) SendMessage(message string) (*Message, error) {
	return handleRequestNow(message, "", "send", ds, nil)
}

func (ds discord) SendEmbed(embed Embed) (*Message, error) {
	return handleRequestNow("", "", "send_embed", ds, &embed)
}

func (ds discord) EditMessage(message, messageID string) (*Message, error) {
	return handleRequestNow(message, messageID, "edit", ds, nil)
}

func (ds discord) EditEmbed(messageID string, embed Embed) (*Message, error) {
	return handleRequestNow("", messageID, "edit_embed", ds, &embed)
}

type DeleteResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (ds discord) DeleteMessage(messageID string) error {
	_, err := handleRequestNow("", messageID, "delete", ds, nil)
	return err
}

func handleRequestNow(message, messageID, todo string, ds discord, embed *Embed) (*Message, error) {
	response, err := sendAPIRequest(message, messageID, todo, ds, embed)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if todo == "delete" {
		if len(resBody) == 0 {
			return nil, nil
		}

		var responseData DeleteResponse
		err = json.Unmarshal(resBody, &responseData)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(responseData.Message)
	}

	var responseData Message
	err = json.Unmarshal(resBody, &responseData)

	return &responseData, err
}

func NewDiscordTextHookService(webhookID, webhookToken string) discordInterfacer {
	return &discord{
		WebhookID:    webhookID,
		WebhookToken: webhookToken,
	}
}
