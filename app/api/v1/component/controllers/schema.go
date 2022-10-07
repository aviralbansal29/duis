package controllers

import (
	"net/http"

	"github.com/aviralbansal29/duis/app/api/v1/component/serializers"
	"github.com/aviralbansal29/duis/app/api/v1/component/services"
	"github.com/labstack/echo/v4"
)

// GetSchema godoc
// @Summary     	Get Schema
// @Description  	Returns schema based on user variant and type
// @Tags          Component
// @Accept        json
// @Produce       json
// @Param         user_id    query   string   true  "User ID"
// @Param         type       query   string   true  "Component Type" Enums(home_page, bottom_bar, recommendation_page)
// @Success       200
// @Failure       400
// @Failure       404
// @Failure       500
// @Router        /api/v1/components [get]
func GetSchema(c echo.Context) error {
	data, err := serializers.SerializeGetSchema(c)
	if err != nil {
		return err
	}

	responseData, err := services.GetSchema(data)
	if err != nil {
		return err
	}

	return c.JSONBlob(http.StatusOK, responseData.Children)
}
