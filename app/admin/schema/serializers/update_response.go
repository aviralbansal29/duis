package serializers

import (
	"net/http"

	"github.com/aviralbansal29/duis/app/models"
	"github.com/aviralbansal29/duis/config"
	"github.com/labstack/echo/v4"
)

// DeserializeUpdateResponse deserializes response for Update API
func DeserializeUpdateResponse(ctx echo.Context, data models.Schema) error {
	var response SchemaRetrieveFormat
	config.DatabaseHandler().Model(&models.Schema{}).First(&response, data.ID)
	return ctx.JSON(http.StatusOK, response)
}
