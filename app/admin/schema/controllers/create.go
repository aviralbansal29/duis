package controllers

import (
	"github.com/aviralbansal29/duis/app/admin/schema/serializers"
	"github.com/aviralbansal29/duis/app/admin/schema/services"
	"github.com/labstack/echo/v4"
)

// Create handles POST API for schema
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
