package routes

import (
	"github.com/aviralbansal29/duis/app/api/v1/user/controllers"
	"github.com/labstack/echo/v4"
)

// AddRoutes adds routes for admin APIs
func AddRoutes(parentRouter *echo.Group) {
	parentRouter.POST("/users", controllers.Register)
}
