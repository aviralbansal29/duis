package controllers

import (
	"github.com/aviralbansal29/duis/app/admin/schema/serializers"
	"github.com/aviralbansal29/duis/app/admin/schema/services"
	"github.com/labstack/echo/v4"
)

// Update handles POST API for schema
func Update(ctx echo.Context) error {
	content, err := serializers.BindUpdateForm(ctx)
	if err != nil {
		return err
	}

	response, err := services.ProcessUpdateData(content, ctx.Param("id"))
	if err != nil {
		return err
	}

	return serializers.DeserializeUpdateResponse(ctx, response)
}
