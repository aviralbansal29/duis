package serializers

import (
	"net/http"

	"github.com/aviralbansal29/duis/app/models"
	"github.com/aviralbansal29/duis/config"
	"github.com/labstack/echo/v4"
)

// DeserializeCreateResponse deserializes response for Create API
func DeserializeCreateResponse(ctx echo.Context, data models.Schema) error {
	var response SchemaRetrieveFormat
	config.DatabaseHandler().Model(&models.Schema{}).First(&response, data.ID)
	return ctx.JSON(http.StatusCreated, response)
}
