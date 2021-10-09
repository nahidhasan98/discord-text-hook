package discord_text_hook

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type workerResponse struct {
	Response *Message
	Error    error
}

func parseError(err error) workerResponse {
	return workerResponse{
		Response: nil,
		Error:    err,
	}
}

func handleWorker(jobs <-chan int, responseChan chan workerResponse, message, messageID, todo string, ds discord) {
	for range jobs {
		response, err := sendAPIRequest(message, messageID, todo, ds)
		if err != nil {
			responseChan <- parseError(err)
			return
		}
		// defer response.Body.Close()

		resBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			responseChan <- parseError(err)
			return
		}

		if todo == "delete" {
			if len(resBody) == 0 {
				responseChan <- parseError(err)
				return
			}

			var responseData DeleteResponse
			err = json.Unmarshal(resBody, &responseData)
			if err != nil {
				responseChan <- parseError(err)
				return
			}

			responseChan <- parseError(errors.New(responseData.Message))
			return
		}

		var responseData Message
		err = json.Unmarshal(resBody, &responseData)

		res := workerResponse{
			Response: &responseData,
			Error:    err,
		}

		responseChan <- res
	}
}
