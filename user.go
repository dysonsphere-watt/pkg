package pkg

import (
	"errors"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dysonsphere-watt/pkg/models"
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

func GetUserIDFromJWT(token string) (int32, error) {
	decodedPayload, err := facades.Auth(http.Background()).Parse(token)
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

func GetUserID(c *app.RequestContext) (int32, error) {
	jwt := GetJWTFromHeader(c)
	userID, err := GetUserIDFromJWT(jwt)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func IsUserAdmin(c *app.RequestContext) bool {
	userID, err := GetUserID(c)
	if err != nil {
		return false
	}

	var roleModel models.Role
	err = facades.Orm().Query().
		Model(&models.User{}).
		Select("role.name").
		Join("JOIN role ON role.id=user.role_id").
		Where(&models.User{ID: userID}).
		FirstOrFail(&roleModel)
	if err != nil {
		return false
	}

	return roleModel.Name == "Admin"
}

func UserHasPermission(c *app.RequestContext, permissionType string) bool {
	userID, err := GetUserID(c)
	if err != nil {
		return false
	}

	var permModel models.Permission
	err = facades.Orm().Query().Model(&models.User{}).Select("permission.*").
		Join("JOIN role r ON user.role_id=r.id").
		Join("LEFT JOIN role_permission rp ON r.id=rp.role_id").
		Join("JOIN permission p ON rp.permission_id=p.id").
		Where(&models.Permission{Type: permissionType}).
		Where(&models.User{ID: userID}).
		FirstOrFail(&permModel)
	if err != nil {
		return false
	}

	return permModel.Type == permissionType
}
