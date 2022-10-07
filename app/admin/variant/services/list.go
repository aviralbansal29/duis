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
)

// ProcessListData processes response for List API
func ProcessListData(ctx echo.Context, params map[string]interface{}) ([]models.Variant, int64, error) {
	var data []models.Variant
	var count int64
	var err error
	query := config.DatabaseHandler().Model(&models.Variant{})
	if val, ok := params["query"]; ok {
		query = query.Where(fmt.Sprintf("name like '%%%s%%'", val))
		delete(params, "query")
	}
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
		log.GetLogger().WithFields(logrus.Fields{"Error": result.Error}).Error("Unable to Find Variants")
		return data, 0, echo.NewHTTPError(http.StatusBadGateway, map[string]interface{}{"errors": result.Error})
	}
	return data, count, nil
}
