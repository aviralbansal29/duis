package serializers

import (
	"net/http"
	"reflect"

	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/aviralbansal29/duis/config"
	"github.com/labstack/echo/v4"
)

// CreateRequestFormat defines format for POST schema request
type CreateRequestFormat struct {
	Name     string      `json:"name" validate:"required"`
	Type     string      `json:"type" validate:"required"`
	Children interface{} `json:"children" validate:"required"`
}

// BindCreateRequest binds data while creating schema
func BindCreateRequest(ctx echo.Context) (CreateRequestFormat, error) {
	var validContent CreateRequestFormat
	if err := (&echo.DefaultBinder{}).BindBody(ctx, &validContent); err != nil {
		return validContent, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": err})
	}
	if err := config.Validator().Struct(validContent); err != nil {
		errorMap := commonUtil.GetErrorMap(reflect.TypeOf(&validContent), err, "json")
		return validContent, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": errorMap})
	}
	return validContent, nil
}
