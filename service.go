package discord_text_hook

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type discordInterfacer interface {
	SendMessage(message string) (*Message, error)
	SendMessageByAnotherProcess(message string) (*Message, error)

	EditMessage(messageID, message string) (*Message, error)
	EditMessageByAnotherProcess(messageID, message string) (*Message, error)

	DeleteMessage(messageID string) error
	DeleteMessageByAnotherProcess(messageID string) error
}

type discord struct {
	WebhookID    string
	WebhookToken string
}

func (ds discord) SendMessage(message string) (*Message, error) {
	return handleRequestNow(message, "", "send", ds)
}

func (ds discord) SendMessageByAnotherProcess(message string) (*Message, error) {
	return handleRequestWithGORoutine(message, "", "send", ds)
}

func (ds discord) EditMessage(message, messageID string) (*Message, error) {
	return handleRequestNow(message, messageID, "edit", ds)
}

func (ds discord) EditMessageByAnotherProcess(messageID, message string) (*Message, error) {
	return handleRequestWithGORoutine(message, messageID, "edit", ds)
}

type DeleteResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (ds discord) DeleteMessage(messageID string) error {
	_, err := handleRequestNow("", messageID, "delete", ds)
	return err
}

func (ds discord) DeleteMessageByAnotherProcess(messageID string) error {
	_, err := handleRequestWithGORoutine("", messageID, "delete", ds)
	return err
}

func handleRequestNow(message, messageID, todo string, ds discord) (*Message, error) {
	response, err := sendAPIRequest(message, messageID, todo, ds)
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

func handleRequestWithGORoutine(message, messageID, todo string, ds discord) (*Message, error) {
	jobs := make(chan int, 5)
	responseChan := make(chan workerResponse, 5)

	go handleWorker(jobs, responseChan, message, messageID, todo, ds)

	jobs <- 1
	defer close(jobs)
	defer close(responseChan)

	response := <-responseChan
	return response.Response, response.Error
}

func NewDiscordHookService(webhookID, webhookToken string) discordInterfacer {
	return &discord{
		WebhookID:    webhookID,
		WebhookToken: webhookToken,
	}
}
