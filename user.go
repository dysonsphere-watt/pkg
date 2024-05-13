package pkg

import (
	"errors"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/http"
)

func GetJWTFromHeader(c *app.RequestContext) string {
	authHeader := string(c.GetHeader("Authorization"))
	spl := strings.Split(authHeader, " ")
	if spl[0] != "Bearer" {
		return ""
	}
	return spl[1]
}

func GetUserIdFromJWT(token string) (int32, error) {
	decodedPayload, err := facades.Auth().Parse(http.Background(), token)
	if err != nil {
		return 0, errors.New("failed to decode JWT")
	}

	userIDStr := decodedPayload.Key
	userID, err := strconv.ParseInt(userIDStr, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(userID), nil
}

func GetUserId(c *app.RequestContext) (int32, error) {
	jwt := GetJWTFromHeader(c)
	userID, err := GetUserIdFromJWT(jwt)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
