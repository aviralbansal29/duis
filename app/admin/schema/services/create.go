package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	serializers "github.com/aviralbansal29/duis/app/admin/schema/serializers"
	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/aviralbansal29/duis/app/models"
	"github.com/aviralbansal29/duis/config"
	"github.com/aviralbansal29/duis/log"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// CreateData Processes Create API data
func CreateData(content serializers.CreateRequestFormat) (models.Schema, error) {
	var data models.Schema
	commonUtil.ConvertType(content, &data)
	jsonSchema, _ := json.Marshal(content.Children)
	data.Children = jsonSchema

	err := validateUniqueName(data, nil)
	if err != nil {
		return data, err
	}

	result := config.DatabaseHandler().Create(&data)
	if result.Error != nil {
		log.GetLogger().WithFields(logrus.Fields{"Error": result.Error}).Error("Unable to Create Schema")
		return data, echo.NewHTTPError(http.StatusBadGateway, map[string]interface{}{"errors": result.Error})
	}
	return data, nil
}

func validateUniqueName(data models.Schema, selfElementID interface{}) error {
	var existingData models.Schema
	filter := config.DatabaseHandler().Where("lower(name) = lower(?)", data.Name)
	if selfElementID != nil {
		filter = filter.Not("id = ?", selfElementID.(int))
	}
	filter.First(&existingData)
	if existingData.ID != 0 {
		return fieldError(
			fmt.Sprintf("Schema with same name already exists. SchemaID : %d", existingData.ID), "name",
		)
	}
	return nil
}

func fieldError(message string, field string) error {
	return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
		"errors": map[string]string{field: message},
	})
}
