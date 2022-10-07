package serializers

import (
	"encoding/json"
	"net/http"

	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/aviralbansal29/duis/app/models"
	"github.com/labstack/echo/v4"
)

// SchemaRetrieveFormat is response format for Retrieve API
type SchemaRetrieveFormat struct {
	ID       uint                     `json:"id"`
	Name     string                   `json:"name"`
	Type     string                   `json:"type"`
	Children []map[string]interface{} `json:"children"`
}

// DeserializeRetrieveResponse deserializers response for Retrieve API
func DeserializeRetrieveResponse(ctx echo.Context, data models.Schema) error {
	var response SchemaRetrieveFormat
	commonUtil.ConvertType(data, &response)
	json.Unmarshal(data.Children, &response.Children)
	return ctx.JSON(http.StatusOK, response)
}
