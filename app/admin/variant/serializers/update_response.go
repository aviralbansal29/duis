package serializers

import (
	"net/http"

	"github.com/aviralbansal29/duis/app/models"
	"github.com/aviralbansal29/duis/config"
	"github.com/labstack/echo/v4"
)

// DeserializeUpdateResponse deserializes response for Update API
func DeserializeUpdateResponse(ctx echo.Context, data models.Variant) error {
	var response VariantRetrieveFormat
	config.DatabaseHandler().Model(&models.Variant{}).First(&response, data.ID)
	return ctx.JSON(http.StatusOK, response)
}
