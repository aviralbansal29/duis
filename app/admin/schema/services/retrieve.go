package services

import (
	"net/http"

	"github.com/aviralbansal29/duis/app/models"
	"github.com/aviralbansal29/duis/config"
	"github.com/aviralbansal29/duis/log"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// ProcessRetrieveData processes data for Retrieve API
func ProcessRetrieveData(id int) (models.Schema, error) {
	var data models.Schema
	result := config.DatabaseHandler().Model(&models.Schema{}).First(&data, id)
	if result.Error != nil {
		log.GetLogger().WithFields(logrus.Fields{"Error": result.Error}).Error("Unable to Find Schema")
		return data, echo.NewHTTPError(http.StatusBadGateway, map[string]interface{}{"errors": result.Error})
	}
	return data, nil
}
