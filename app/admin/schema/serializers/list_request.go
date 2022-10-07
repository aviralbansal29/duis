package serializers

import (
	"net/http"

	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/aviralbansal29/duis/log"
	"github.com/labstack/echo/v4"
)

type listRequestParam struct {
	ID        []string `query:"id" json:"id" validate:"omitempty"`
	Name      string   `query:"name" json:"name" validate:"omitempty"`
	Type      string   `query:"query" json:"type" validate:"omitempty"`
	VariantID uint     `query:"variant_id" json:"variant_id" validate:"omitempty"`
}

// BindListParams binds params for List API
func BindListParams(ctx echo.Context) (map[string]interface{}, error) {
	var filters listRequestParam
	filterMap := map[string]interface{}{}
	if err := (&echo.DefaultBinder{}).BindQueryParams(ctx, &filters); err != nil {
		return filterMap, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": err})
	}
	commonUtil.ConvertType(filters, &filterMap)
	var shouldDelete bool
	for key, value := range filterMap {
		switch v := value.(type) {
		case string:
			shouldDelete = value.(string) == ""
		case float64:
			shouldDelete = value.(float64) == 0
		case nil:
			shouldDelete = true
		default:
			shouldDelete = true
			log.GetLogger().WithField("Param Type", v).Error("Unknown param type")
		}
		if shouldDelete {
			delete(filterMap, key)
		}
	}
	return filterMap, nil
}
