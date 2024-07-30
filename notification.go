package pkg

import (
	"errors"
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/goravel/framework/facades"
)

type PushTopicBody struct {
	Topic        string            `json:"topic"`
	TemplateCode string            `json:"template_code"`
	SharedKey    string            `json:"shared_key"`
	Data         map[string]string `json:"data_str"`
	TemplateMap  map[string]string `json:"template_map_str"`
}

type PushTokensBody struct {
	Tokens       []string          `json:"tokens"`
	TemplateCode string            `json:"template_code"`
	SharedKey    string            `json:"shared_key"`
	Data         map[string]string `json:"data_str"`
	TemplateMap  map[string]string `json:"template_map_str"`
}

type SendSMSBody struct {
	TemplateCode string            `json:"template_code"`
	TemplateMap  map[string]string `json:"template_map"`
	Numbers      []string          `json:"numbers"`
	SharedKey    string            `json:"shared_key"`
}

type SendEmailBody struct {
	TemplateCode string            `json:"template_code"`
	TemplateMap  map[string]string `json:"template_map"`
	Emails       []string          `json:"emails"`
	SharedKey    string            `json:"shared_key"`
}

type PushNotificationResponse struct {
	StatusCode int32  `json:"status_code"`
	Msg        string `json:"msg"`
}

// Send push notifications to topic
func SendPushNotificationTopic(identifier, templateCode string, data map[string]string, templateMap map[string]string) error {
	var resBody PushNotificationResponse
	var err error

	url := facades.Config().GetString("WATT_NOTIFICATION_PUSH_TOPIC_URL", "")
	if url == "" {
		return errors.New("WATT_NOTIFICATION_PUSH_TOPIC_URL is not set, unable to send push notification")
	}

	reqBody := PushTopicBody{
		Topic:        identifier,
		TemplateCode: templateCode,
		SharedKey:    facades.Config().GetString("WATT_NOTIFICATION_SHARED_KEY"),
		Data:         data,
		TemplateMap:  templateMap,
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
func SendPushNotificationTokens(tokens []string, templateCode string, data map[string]string, templateMap map[string]string) error {
	var resBody PushNotificationResponse
	var err error

	url := facades.Config().GetString("WATT_NOTIFICATION_PUSH_TOKENS_URL", "")
	if url == "" {
		return errors.New("WATT_NOTIFICATION_PUSH_TOKENS_URL is not set, unable to send push notification ")
	}

	reqBody := PushTokensBody{
		Tokens:       tokens,
		TemplateCode: templateCode,
		SharedKey:    facades.Config().GetString("WATT_NOTIFICATION_SHARED_KEY"),
		Data:         data,
		TemplateMap:  templateMap,
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

func SendSMS(numbers []string, templateCode string, templateData map[string]string) error {
	var resBody PushNotificationResponse
	var err error

	url := facades.Config().GetString("WATT_NOTIFICATION_SEND_SMS_URL", "")
	if url == "" {
		return errors.New("WATT_NOTIFICATION_SEND_SMS_URL is not set, unable to send sms")
	}

	reqBody := SendSMSBody{
		Numbers:      numbers,
		TemplateCode: templateCode,
		SharedKey:    facades.Config().GetString("WATT_NOTIFICATION_SHARED_KEY"),
		TemplateMap:  templateData,
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

func SendEmails(emails []string, templateCode string, templateData map[string]string) error {
	var resBody PushNotificationResponse
	var err error

	url := facades.Config().GetString("WATT_NOTIFICATION_SEND_EMAIL_URL", "")
	if url == "" {
		return errors.New("WATT_NOTIFICATION_SEND_EMAIL_URL is not set, unable to send email")
	}

	reqBody := SendEmailBody{
		Emails:       emails,
		TemplateCode: templateCode,
		SharedKey:    facades.Config().GetString("WATT_NOTIFICATION_SHARED_KEY"),
		TemplateMap:  templateData,
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
