package controllers

import (
	"net/http"

	"github.com/aviralbansal29/duis/app/admin/schema/serializers"
	"github.com/aviralbansal29/duis/app/models"
	"github.com/aviralbansal29/duis/config"
	"github.com/labstack/echo/v4"
)

// Delete handles DELETE request
func Delete(ctx echo.Context) error {
	id, err := serializers.BindRetrieveParam(ctx)
	if err != nil {
		return err
	}

	schema := models.Schema{}
	config.DatabaseHandler().First(&schema, id)

	config.DatabaseHandler().Delete(&schema)

	return ctx.JSON(http.StatusOK, schema)
}
