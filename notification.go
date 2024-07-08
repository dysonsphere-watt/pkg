package pkg

import (
	"encoding/json"
	"errors"
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/goravel/framework/facades"
)

type PushTopicBody struct {
	Topic        string `json:"topic"`
	TemplateCode string `json:"template_code"`
	SharedKey    string `json:"shared_key"`
	DataStr      string `json:"data_str"`
}

type PushTokensBody struct {
	Tokens       []string `json:"tokens"`
	TemplateCode string   `json:"template_code"`
	SharedKey    string   `json:"shared_key"`
	DataStr      string   `json:"data_str"`
}

type PushNotificationResponse struct {
	StatusCode int32  `json:"status_code"`
	Msg        string `json:"msg"`
}

// Send push notifications to topic
func SendPushNotificationTopic(identifier, templateCode string, data *map[string]string) error {
	var resBody PushNotificationResponse

	url := facades.Config().GetString("WATT_NOTIFICATION_PUSH_TOPIC_URL", "")
	if url == "" {
		return errors.New("WATT_NOTIFICATION_PUSH_TOPIC_URL is not set, unable to send push notification")
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return errors.New("error converting \"data\" into a JSON string")
	}

	reqBody := PushTopicBody{
		Topic:        identifier,
		TemplateCode: templateCode,
		SharedKey:    facades.Config().GetString("WATT_NOTIFICATION_SHARED_KEY"),
		DataStr:      string(dataBytes),
	}

	client := resty.New()
	res, err := client.R().
		SetBody(reqBody).
		SetResult(&resBody).
		SetError(&resBody).
		Post(url)
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("error pushing notifications: %s", resBody.Msg)
	}

	return nil
}

// Send push notifications to a bunch of tokens
func SendPushNotificationTokens(tokens []string, templateCode string, data *map[string]string) error {
	var resBody PushNotificationResponse

	url := facades.Config().GetString("WATT_NOTIFICATION_PUSH_TOKENS_URL", "")
	if url == "" {
		return errors.New("WATT_NOTIFICATION_PUSH_TOKENS_URL is not set, unable to send push notification ")
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return errors.New("error converting \"data\" into a JSON string")
	}

	reqBody := PushTokensBody{
		Tokens:       tokens,
		TemplateCode: templateCode,
		SharedKey:    facades.Config().GetString("WATT_NOTIFICATION_SHARED_KEY"),
		DataStr:      string(dataBytes),
	}

	client := resty.New()
	res, err := client.R().
		SetBody(reqBody).
		SetResult(&resBody).
		SetError(&resBody).
		Post(url)
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("error pushing notifications: %s", resBody.Msg)
	}

	return nil
}
