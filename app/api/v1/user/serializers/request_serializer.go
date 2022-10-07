package serializers

import (
	"net/http"
	"reflect"

	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/aviralbansal29/duis/config"
	"github.com/labstack/echo/v4"
)

// RegisterUserRequestBody defines format for Registering User
type RegisterUserRequestBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// SerializeRegisterUser serializes the request body
func SerializeRegisterUser(ctx echo.Context) (RegisterUserRequestBody, error) {
	var validBody RegisterUserRequestBody
	if err := (&echo.DefaultBinder{}).BindBody(ctx, &validBody); err != nil {
		return validBody, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": err})
	}
	if err := config.Validator().Struct(validBody); err != nil {
		errorMap := commonUtil.GetErrorMap(reflect.TypeOf(&validBody), err, "json")
		return validBody, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": errorMap})
	}
	return validBody, nil
}
