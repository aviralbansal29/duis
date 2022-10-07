package middleware

import (
	"net/http"
	"reflect"
	"strings"

	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/aviralbansal29/duis/config"
	"github.com/labstack/echo/v4"
)

type adminAuthHeader struct {
	APIKey string `header:"Authorization" binding:"required"`
}

// AdminAuthorize authorizes admin APIs
func AdminAuthorize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// fetch user id and user token from headers
		authHeader := adminAuthHeader{}
		if err := (&echo.DefaultBinder{}).BindHeaders(c, &authHeader); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": err})
		}
		if err := config.Validator().Struct(authHeader); err != nil {
			errorMap := commonUtil.GetErrorMap(reflect.TypeOf(&authHeader), err, "header")
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": errorMap})
		}
		if len(strings.Split(authHeader.APIKey, "Bearer ")) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return next(c)
	}
}
