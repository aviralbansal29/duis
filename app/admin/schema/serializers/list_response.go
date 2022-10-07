package serializers

import (
	"fmt"
	"net/http"

	commonUtil "github.com/aviralbansal29/duis/app/common/utils"
	"github.com/aviralbansal29/duis/app/models"
	"github.com/labstack/echo/v4"
)

// SchemaListResponseFormat defines format for List API response
type SchemaListResponseFormat struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// DeserializeListResponse deserializes response for List API
func DeserializeListResponse(ctx echo.Context, data []models.Schema, count int64) error {
	ctx.Response().Header().Set("X-Total-Count", fmt.Sprintf("%d", count))
	ctx.Response().Header().Set("Access-Control-Expose-Headers", "X-Total-Count")
	var response []SchemaListResponseFormat
	commonUtil.ConvertType(data, &response)
	return ctx.JSON(http.StatusOK, response)
}
