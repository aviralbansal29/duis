package serializers

import (
	"net/http"
	"reflect"

	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/aviralbansal29/duis/config"
	"github.com/labstack/echo/v4"
)

type getSchemaParam struct {
	Type   string `query:"type" validate:"required,oneof=home_page bottom_bar recommendation_page"`
	UserID string `query:"user_id" validate:"required"`
}

func SerializeGetSchema(ctx echo.Context) (map[string]string, error) {
	var validParams getSchemaParam

	if err := (&echo.DefaultBinder{}).BindQueryParams(ctx, &validParams); err != nil {
		return map[string]string{}, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": err})
	}
	if err := config.Validator().Struct(validParams); err != nil {
		errorMap := commonUtil.GetErrorMap(reflect.TypeOf(&validParams), err)
		return map[string]string{}, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": errorMap})
	}
	response := map[string]string{
		"type":    validParams.Type,
		"user_id": validParams.UserID,
	}
	return response, nil
}
