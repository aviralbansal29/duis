package serializers

import (
	"net/http"
	"reflect"

	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/aviralbansal29/duis/config"
	"github.com/labstack/echo/v4"
)

// SchemaUpdateFormFormat defines form data for Update API
type SchemaUpdateFormFormat struct {
	Name     string      `json:"name" validate:"required"`
	Type     string      `json:"type" validate:"required"`
	Children interface{} `json:"children" validate:"required"`
}

// BindUpdateForm serializes form data for Update API
func BindUpdateForm(ctx echo.Context) (SchemaUpdateFormFormat, error) {
	var validContent SchemaUpdateFormFormat
	if err := (&echo.DefaultBinder{}).BindBody(ctx, &validContent); err != nil {
		return validContent, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": err})
	}
	if err := config.Validator().Struct(validContent); err != nil {
		errorMap := commonUtil.GetErrorMap(reflect.TypeOf(&validContent), err, "json")
		return validContent, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": errorMap})
	}
	return validContent, nil
}
