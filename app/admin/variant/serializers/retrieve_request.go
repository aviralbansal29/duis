package serializers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// BindRetrieveParam serializes form data for Create API
func BindRetrieveParam(ctx echo.Context) (int, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return id, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": "Invalid ID"})
	}
	return id, nil
}
