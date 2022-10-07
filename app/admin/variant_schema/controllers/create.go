package controllers

import (
	"net/http"

	"github.com/aviralbansal29/duis/app/admin/variant_schema/serializers"
	"github.com/aviralbansal29/duis/app/admin/variant_schema/services"
	"github.com/labstack/echo/v4"
)

// Create handles POST API for variant_schema
func Create(ctx echo.Context) error {
	content, err := serializers.BindCreateRequest(ctx)
	if err != nil {
		return err
	}

	response, err := services.CreateData(content)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, response)
}
