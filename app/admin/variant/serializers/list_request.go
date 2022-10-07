package serializers

import (
	"net/http"

	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/labstack/echo/v4"
)

type listRequestParam struct {
	ID    []string `query:"id" json:"id" validate:"omitempty"`
	Name  string   `query:"name" json:"name" validate:"omitempty"`
	Query string   `query:"query" json:"query"`
}

// BindListParams binds params for List API
func BindListParams(ctx echo.Context) (map[string]interface{}, error) {
	var filters listRequestParam
	filterMap := map[string]interface{}{}
	if err := (&echo.DefaultBinder{}).BindQueryParams(ctx, &filters); err != nil {
		return filterMap, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": err})
	}
	commonUtil.ConvertType(filters, &filterMap)
	for key, value := range filterMap {
		if value == "" || value == nil {
			delete(filterMap, key)
		}
	}
	return filterMap, nil
}
