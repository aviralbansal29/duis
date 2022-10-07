package serializers

import (
	"net/http"

	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/aviralbansal29/duis/app/models"
	"github.com/labstack/echo/v4"
)

// VariantRetrieveFormat is response format for Retrieve API
type VariantRetrieveFormat struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DeserializeRetrieveResponse deserializers response for Retrieve API
func DeserializeRetrieveResponse(ctx echo.Context, data models.Variant) error {
	var response VariantRetrieveFormat
	commonUtil.ConvertType(data, &response)
	return ctx.JSON(http.StatusOK, response)
}
