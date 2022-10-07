package services

import (
	"net/http"

	"github.com/aviralbansal29/duis/app/models"
	"github.com/aviralbansal29/duis/config"
	"github.com/labstack/echo/v4"
)

func GetSchema(data map[string]string) (models.Schema, error) {
	var schema models.Schema
	var user models.User
	config.DatabaseHandler().First(&user, data["user_id"])
	if user.ID == 0 {
		return schema, echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{"errors": map[string]string{"user_id": "User Not Found"}})
	}

	err := config.DatabaseHandler().Table("schemas").Joins(
		"left join schema_mappings on schemas.id = schema_mappings.schema_id",
	).Where("schemas.type = ? and schema_mappings.variant_id = ?", data["type"], user.VariantID).First(&schema).Error

	return schema, err
}
