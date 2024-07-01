package pkg

import (
	"errors"
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/goravel/framework/facades"
)

type PushNotificationBody struct {
	Type      int32  `json:"type"`
	Token     string `json:"string"`
	Topic     string `json:"topic"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	ImageURL  string `json:"image_url"`
	SharedKey string `json:"shared_key"`
}

type PushNotificationResponse struct {
	StatusCode int32  `json:"status_code"`
	Msg        string `json:"msg"`
}

func SendPushNotification(pushType int32, identifier, title, body, imageURL string) error {
	var resBody PushNotificationResponse

	url := facades.Config().GetString("WATT_NOTIFICATION_PUSH_URL", "")
	if url == "" {
		return errors.New("WATT_NOTIFICATION_PUSH_URL is not set, unable to send push notification")
	}

	reqBody := PushNotificationBody{
		Type:      pushType,
		Title:     title,
		Body:      body,
		ImageURL:  imageURL,
		SharedKey: facades.Config().GetString("WATT_NOTIFICATION_SHARED_KEY"),
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
