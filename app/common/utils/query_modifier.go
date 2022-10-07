package commonUtil

import (
	"net/http"
	"strconv"

	"github.com/aviralbansal29/duis/log"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// QueyOffsetAndLimit adds oddset and limit to query based on request params
func QueyOffsetAndLimit(ctx echo.Context, query *gorm.DB) (*gorm.DB, error) {
	offset, end := ctx.QueryParam("_start"), ctx.QueryParam("_end")
	if offset != "" && end != "" {
		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			log.GetLogger().WithFields(logrus.Fields{"Error": err.Error()}).Error("Unable to change offset to int")
			return query, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": err.Error()})
		}
		endInt, err := strconv.Atoi(end)
		if err != nil {
			log.GetLogger().WithFields(logrus.Fields{"Error": err.Error()}).Error("Unable to change limit to int")
			return query, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": err.Error()})
		}
		query = query.Offset(offsetInt).Limit(endInt - offsetInt)
	}
	return query, nil
}

// QuerySortKeyAndOrder adds sort and order to query based on request params
func QuerySortKeyAndOrder(ctx echo.Context, query *gorm.DB) (*gorm.DB, error) {
	sortKey, order := ctx.QueryParam("_sort"), ctx.QueryParam("_order")
	if sortKey != "" {
		if order != "" {
			sortKey = sortKey + " " + order
		}
		query = query.Order(sortKey)
	}
	return query, nil
}
