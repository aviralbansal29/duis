package services

import (
	"fmt"
	"net/http"

	serializers "github.com/aviralbansal29/duis/app/admin/variant_schema/serializers"
	"github.com/aviralbansal29/duis/app/models"
	"github.com/aviralbansal29/duis/config"
	"github.com/aviralbansal29/duis/log"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// CreateData Processes Create API data
func CreateData(content serializers.CreateRequestFormat) (map[string]interface{}, error) {
	data := map[string]interface{}{"variant_id": content.VariantID, "schema_id": content.SchemaID}

	err := validateUniqueSchemaType(data)
	if err != nil {
		return data, err
	}

	result := config.DatabaseHandler().Table("schema_mappings").Create(&data)
	if result.Error != nil {
		log.GetLogger().WithFields(logrus.Fields{"Error": result.Error}).Error("Unable to Create Link")
		return data, echo.NewHTTPError(http.StatusBadGateway, map[string]interface{}{"errors": result.Error})
	}
	data["id"] = 1
	return data, nil
}

func validateUniqueSchemaType(data map[string]interface{}) error {
	var schemaType string
	config.DatabaseHandler().Model(&models.Schema{}).Select("type").First(&schemaType, data["schema_id"])
	var count int64
	config.DatabaseHandler().Table("schema_mappings").Joins(
		"left join schemas on schemas.id = schema_mappings.schema_id",
	).Where("schemas.type = ?", schemaType).Count(&count)
	if count != 0 {
		return fieldError(
			fmt.Sprintf("Variant already has a %s", schemaType), "variant_id",
		)
	}
	return nil
}

func fieldError(message string, field string) error {
	return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
		"errors": map[string]string{field: message},
	})
}
