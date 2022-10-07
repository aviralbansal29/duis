package controllers

import (
	"net/http"

	"github.com/aviralbansal29/duis/app/admin/variant/serializers"
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

	variant := models.Variant{}
	config.DatabaseHandler().First(&variant, id)

	config.DatabaseHandler().Delete(&variant)

	return ctx.JSON(http.StatusOK, variant)
}
