package serializers

import (
	"net/http"
	"reflect"

	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/aviralbansal29/duis/config"
	"github.com/labstack/echo/v4"
)

// VariantUpdateFormFormat defines form data for Update API
type VariantUpdateFormFormat struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

// BindUpdateForm serializes form data for Update API
func BindUpdateForm(ctx echo.Context) (VariantUpdateFormFormat, error) {
	var validContent VariantUpdateFormFormat
	if err := (&echo.DefaultBinder{}).BindBody(ctx, &validContent); err != nil {
		return validContent, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": err})
	}
	if err := config.Validator().Struct(validContent); err != nil {
		errorMap := commonUtil.GetErrorMap(reflect.TypeOf(&validContent), err, "json")
		return validContent, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": errorMap})
	}
	return validContent, nil
}
