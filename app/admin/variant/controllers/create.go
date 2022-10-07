package controllers

import (
	"github.com/aviralbansal29/duis/app/admin/variant/serializers"
	"github.com/aviralbansal29/duis/app/admin/variant/services"
	"github.com/labstack/echo/v4"
)

// Create handles POST API for variants
func Create(ctx echo.Context) error {
	content, err := serializers.BindCreateRequest(ctx)
	if err != nil {
		return err
	}

	response, err := services.CreateData(content)
	if err != nil {
		return err
	}

	return serializers.DeserializeCreateResponse(ctx, response)
}
