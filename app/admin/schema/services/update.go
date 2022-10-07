package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aviralbansal29/duis/app/admin/schema/serializers"
	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/aviralbansal29/duis/app/models"
	"github.com/aviralbansal29/duis/config"
	"github.com/aviralbansal29/duis/log"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// ProcessUpdateData processes data for Update API
func ProcessUpdateData(content serializers.SchemaUpdateFormFormat, idParam string) (models.Schema, error) {
	id, _ := strconv.Atoi(idParam)
	var schema models.Schema
	result := config.DatabaseHandler().Find(&schema, id)
	if result.Error != nil {
		return schema, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": result.Error})
	}
	var data models.Schema
	commonUtil.ConvertType(content, &data)
	data.Children, _ = json.Marshal(content.Children)

	err := validateUniqueName(data, id)
	if err != nil {
		return schema, err
	}

	result = config.DatabaseHandler().Model(&schema).Updates(data)
	if result.Error != nil {
		log.GetLogger().WithFields(logrus.Fields{"Error": result.Error}).Error("Unable to Update Schema")
		return schema, echo.NewHTTPError(http.StatusBadGateway, map[string]interface{}{"errors": result.Error})
	}
	return schema, nil
}
