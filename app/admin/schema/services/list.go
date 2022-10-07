package services

import (
	"fmt"
	"net/http"

	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/aviralbansal29/duis/app/models"
	"github.com/aviralbansal29/duis/config"
	"github.com/aviralbansal29/duis/log"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// ProcessListData processes response for List API
func ProcessListData(ctx echo.Context, params map[string]interface{}) ([]models.Schema, int64, error) {
	var data []models.Schema
	var count int64
	var err error
	query := config.DatabaseHandler().Model(&models.Schema{})
	if val, ok := params["query"]; ok {
		query = query.Where(fmt.Sprintf("name like '%%%s%%'", val))
		delete(params, "query")
	}
	query, params = applyVariantFilter(query, params)
	query = query.Where(params)
	query.Count(&count)
	query, err = commonUtil.QueyOffsetAndLimit(ctx, query)
	if err != nil {
		return data, 0, err
	}
	query, err = commonUtil.QuerySortKeyAndOrder(ctx, query)
	if err != nil {
		return data, 0, err
	}
	result := query.Find(&data)
	if result.Error != nil {
		log.GetLogger().WithFields(logrus.Fields{"Error": result.Error}).Error("Unable to Find Schemas")
		return data, 0, echo.NewHTTPError(http.StatusBadGateway, map[string]interface{}{"errors": result.Error})
	}
	return data, count, nil
}

func applyVariantFilter(query *gorm.DB, params map[string]interface{}) (*gorm.DB, map[string]interface{}) {
	if variantID, found := params["variant_id"]; found {
		query = query.Joins("left join schema_mappings on schemas.id = schema_mappings.schema_id").Where(
			"schema_mappings.variant_id = ?", variantID,
		)
		delete(params, "variant_id")
	}
	return query, params
}
