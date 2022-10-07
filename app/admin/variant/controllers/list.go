package controllers

import (
	"github.com/aviralbansal29/duis/app/admin/variant/serializers"
	"github.com/aviralbansal29/duis/app/admin/variant/services"
	"github.com/labstack/echo/v4"
)

// List returns list of variants
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
