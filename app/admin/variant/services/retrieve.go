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
func ProcessRetrieveData(id int) (models.Variant, error) {
	var data models.Variant
	result := config.DatabaseHandler().Model(&models.Variant{}).First(&data, id)
	if result.Error != nil {
		log.GetLogger().WithFields(logrus.Fields{"Error": result.Error}).Error("Unable to Find Element")
		return data, echo.NewHTTPError(http.StatusBadGateway, map[string]interface{}{"errors": result.Error})
	}
	return data, nil
}
