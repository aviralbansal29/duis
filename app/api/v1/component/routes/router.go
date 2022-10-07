package routes

import (
	controllers "github.com/aviralbansal29/duis/app/api/v1/component/controllers"
	"github.com/labstack/echo/v4"
)

// AddRoutes adds routes for admin APIs
func AddRoutes(parentRouter *echo.Group) {
	parentRouter.GET("/components", controllers.GetSchema)
}
