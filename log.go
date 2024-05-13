package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goravel/framework/facades"
)

type LogType string

const (
	LogTypeInfo  LogType = "INFO"
	LogTypeDebug LogType = "DEBUG"
	LogTypeWarn  LogType = "WARN"
	LogTypeError LogType = "ERROR"
	LogTypeFatal LogType = "FATAL"
)

type LogBody struct {
	ApplicationType  string `json:"application_type"`
	LogType          string `json:"log_type"`
	DetectedIP       string `json:"detected_ip"`
	DetectedPlatform string `json:"detected_platform"`
	Content          string `json:"content"`
	UserID           int    `json:"user_id"`
	RobotID          int    `json:"robot_id"`
	BookingID        int    `json:"booking_id"`
	OrderID          int    `json:"order_id"`
}

// Logs both locally and to a central log server
func DistLog(logType LogType, detectedIP, detectedPlatform, content string, userID, robotID, bookingID, orderID int) {
	applicationType := facades.Config().GetString("APP_MODULE", "Watt-Generic")

	body := LogBody{
		ApplicationType:  applicationType,
		LogType:          string(logType),
		DetectedIP:       detectedIP,
		DetectedPlatform: detectedPlatform,
		Content:          content,
		UserID:           userID,
		RobotID:          robotID,
		BookingID:        bookingID,
		OrderID:          orderID,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error marshaling JSON body:", err.Error())
		return
	}

	req, err := http.NewRequest(http.MethodPost, facades.Config().GetString("WATT_LOG_CREATE_URL"), bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating HTTP request for logging:", err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	if logType == LogTypeInfo {
		facades.Log().Info(applicationType, detectedIP, detectedPlatform, content)
	} else if logType == LogTypeDebug {
		facades.Log().Debug(applicationType, detectedIP, detectedPlatform, content)
	} else if logType == LogTypeWarn {
		facades.Log().Warning(applicationType, detectedIP, detectedPlatform, content)
	} else if logType == LogTypeError {
		facades.Log().Error(applicationType, detectedIP, detectedPlatform, content)
	} else if logType == LogTypeFatal {
		facades.Log().Fatal(applicationType, detectedIP, detectedPlatform, content)
	}
}
