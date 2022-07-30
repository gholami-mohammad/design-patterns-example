package notification

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

func Send(message string, recipient string) error {
	apiKey := os.Getenv("notification.api_key")
	apiToken := os.Getenv("notification.token")
	notificationServer := os.Getenv("notification.server")
	apiPath := notificationServer + "/api/send"

	requestBody := sendNotificationRequestBody{
		Message:   message,
		Recipient: recipient,
	}

	bts, err := json.Marshal(&requestBody)
	if err != nil {
		return err
	}

	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, apiPath, bytes.NewReader(bts))
	if err != nil {
		return err
	}
	req.Header.Set("api_key", apiKey)
	req.Header.Set("api_token", apiToken)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New(string(response))
	}

	return nil
}

type sendNotificationRequestBody struct {
	Message   string
	Recipient string
}
