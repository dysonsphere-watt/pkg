package pkg

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/goravel/framework/facades"
)

func GetUserID(c *app.RequestContext) (string, error) {
	userIDEnc := string(c.GetHeader(UserIDHeaderKey))
	userIDBytes, err := Decrypt([]byte(facades.Config().GetString("WATT_AES_SHARED_KEY", "")), userIDEnc)
	if err != nil {
		return "", err
	}

	userID := string(userIDBytes)
	return userID, nil
}
