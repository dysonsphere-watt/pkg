package pkg

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	resty "github.com/go-resty/resty/v2"
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
	Method           string `json:"method"`
	Path             string `json:"path"`
	RequestBody      string `json:"request_body"`
	Content          string `json:"content"`
	UserID           int32  `json:"user_id"`
	RobotID          int32  `json:"robot_id"`
	BookingID        int32  `json:"booking_id"`
	OrderID          int32  `json:"order_id"`
}

// Logs both locally and to a distributed log server
func DistLog(logType LogType, c *app.RequestContext, reqBody, content string, userID, robotID, bookingID, orderID int32) {
	if c == nil {
		c = &app.RequestContext{}
	}

	cUserID, err := GetUserID(c)
	if err == nil && userID == 0 {
		userID = cUserID
	}

	applicationType := facades.Config().GetString("APP_MODULE", "Watt-Generic")
	detectedIP := c.ClientIP()
	detectedPlatform := string(c.UserAgent())
	method := string(c.Method())
	path := string(c.Path())

	prettyPrint := fmt.Sprintf("[%s] IP: %s, User-Agent: %s: %s", applicationType, detectedIP, detectedPlatform, content)

	switch logType {
	case LogTypeInfo:
		facades.Log().Info(prettyPrint)
	case LogTypeDebug:
		facades.Log().Debug(prettyPrint)
	case LogTypeWarn:
		facades.Log().Warning(prettyPrint)
	case LogTypeError:
		facades.Log().Error(prettyPrint)
	case LogTypeFatal:
		facades.Log().Fatal(prettyPrint)
	default:
		facades.Log().Info(prettyPrint)
	}

	url := facades.Config().GetString("WATT_LOG_CREATE_URL", "")
	if url == "" {
		facades.Log().Error("WATT_LOG_CREATE_URL is not set, unable to log to the server!")
		return
	}

	logBody := LogBody{
		ApplicationType:  applicationType,
		LogType:          string(logType),
		DetectedIP:       detectedIP,
		DetectedPlatform: detectedPlatform,
		Method:           method,
		Path:             path,
		RequestBody:      reqBody,
		Content:          content,
		UserID:           userID,
		RobotID:          robotID,
		BookingID:        bookingID,
		OrderID:          orderID,
	}

	client := resty.New()
	client.R().SetBody(logBody).Post(url)
}
