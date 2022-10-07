package controllers

import (
	"github.com/aviralbansal29/duis/app/admin/variant/serializers"
	"github.com/aviralbansal29/duis/app/admin/variant/services"
	"github.com/labstack/echo/v4"
)

// Retrieve handles Retrieve API
func Retrieve(ctx echo.Context) error {
	id, err := serializers.BindRetrieveParam(ctx)
	if err != nil {
		return err
	}

	response, err := services.ProcessRetrieveData(id)
	if err != nil {
		return err
	}

	return serializers.DeserializeRetrieveResponse(ctx, response)
}
