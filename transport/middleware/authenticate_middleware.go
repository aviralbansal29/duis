package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/aviralbansal29/duis/config"
	"github.com/labstack/echo/v4"
)

type authenticationHeader struct {
	TnlUserID  string `header:"X-TNL-USER-ID" binding:"required,number"`
	TnlTokenID string `header:"X-TNL-TOKEN" binding:"required"`
	AppID      string `header:"X-TNL-APPID"`
}

// Authenticate authenticates the API using TLLMS API
func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// fetch user id and user token from headers
		authHeader := authenticationHeader{}
		if err := (&echo.DefaultBinder{}).BindHeaders(c, &authHeader); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": err})
		}
		if err := config.Validator().Struct(authHeader); err != nil {
			errorMap := commonUtil.GetErrorMap(reflect.TypeOf(&authHeader), err, "header")
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": errorMap})
		}
		env := config.GetEnv()
		resp, err := config.RestyClient().R().
			SetQueryParams(map[string]string{
				"token":  authHeader.TnlTokenID,
				"key":    env.GetString("tllms.client_id"),
				"secret": env.GetString("tllms.client_secret"),
			}).
			SetHeader("X-TNL-APPID", authHeader.AppID).
			Get(fmt.Sprintf(
				"%s/%s/%s",
				env.GetString("tllms.host"),
				env.GetString("tllms.authenticate_endpoint"),
				authHeader.TnlUserID,
			))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, errors.New("Unexpected error Occured"))
		} else if resp.StatusCode() != 200 {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return next(c)
	}
}
