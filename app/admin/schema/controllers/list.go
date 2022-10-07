package controllers

import (
	"github.com/aviralbansal29/duis/app/admin/schema/serializers"
	"github.com/aviralbansal29/duis/app/admin/schema/services"
	"github.com/labstack/echo/v4"
)

// List returns list of schema
func List(ctx echo.Context) error {
	params, err := serializers.BindListParams(ctx)
	if err != nil {
		return err
	}

	data, count, err := services.ProcessListData(ctx, params)
	if err != nil {
		return err
	}

	return serializers.DeserializeListResponse(ctx, data, count)
}
