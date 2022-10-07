package services

import (
	"fmt"
	"net/http"

	serializers "github.com/aviralbansal29/duis/app/admin/variant/serializers"
	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/aviralbansal29/duis/app/models"
	"github.com/aviralbansal29/duis/config"
	"github.com/aviralbansal29/duis/log"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// CreateData Processes Create API data
func CreateData(content serializers.CreateRequestFormat) (models.Variant, error) {
	var data models.Variant
	commonUtil.ConvertType(content, &data)

	err := validateUniqueName(data, nil)
	if err != nil {
		return data, err
	}

	result := config.DatabaseHandler().Create(&data)
	if result.Error != nil {
		log.GetLogger().WithFields(logrus.Fields{"Error": result.Error}).Error("Unable to Create Variant")
		return data, echo.NewHTTPError(http.StatusBadGateway, map[string]interface{}{"errors": result.Error})
	}
	return data, nil
}

func validateUniqueName(data models.Variant, selfElementID interface{}) error {
	var existingData models.Variant
	filter := config.DatabaseHandler().Where("lower(name) = lower(?)", data.Name)
	if selfElementID != nil {
		filter = filter.Not("id = ?", selfElementID.(int))
	}
	filter.First(&existingData)
	if existingData.ID != 0 {
		return fieldError(
			fmt.Sprintf("Variant with same name already exists. VariantID : %d", existingData.ID), "name",
		)
	}
	return nil
}

func fieldError(message string, field string) error {
	return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
		"errors": map[string]string{field: message},
	})
}
