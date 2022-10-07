package services

import (
	"net/http"
	"strconv"

	"github.com/aviralbansal29/duis/app/admin/variant/serializers"
	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/aviralbansal29/duis/app/models"
	"github.com/aviralbansal29/duis/config"
	"github.com/aviralbansal29/duis/log"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// ProcessUpdateData processes data for Update API
func ProcessUpdateData(content serializers.VariantUpdateFormFormat, idParam string) (models.Variant, error) {
	id, _ := strconv.Atoi(idParam)
	var variant models.Variant
	result := config.DatabaseHandler().Find(&variant, id)
	if result.Error != nil {
		return variant, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": result.Error})
	}
	var data models.Variant
	commonUtil.ConvertType(content, &data)

	err := validateUniqueName(data, id)
	if err != nil {
		return variant, err
	}

	result = config.DatabaseHandler().Model(&variant).Updates(data)
	if result.Error != nil {
		log.GetLogger().WithFields(logrus.Fields{"Error": result.Error}).Error("Unable to Update Variant")
		return variant, echo.NewHTTPError(http.StatusBadGateway, map[string]interface{}{"errors": result.Error})
	}
	return variant, nil
}
