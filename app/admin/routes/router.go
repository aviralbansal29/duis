package routes

import (
	schemaControllers "github.com/aviralbansal29/duis/app/admin/schema/controllers"
	variantControllers "github.com/aviralbansal29/duis/app/admin/variant/controllers"
	variantSchemaControllers "github.com/aviralbansal29/duis/app/admin/variant_schema/controllers"
	"github.com/labstack/echo/v4"
)

// AddRoutes adds routes for admin APIs
func AddRoutes(parentRouter *echo.Group) {
	parentRouter.POST("/variants", variantControllers.Create)
	parentRouter.GET("/variants", variantControllers.List)
	parentRouter.GET("/variants/:id", variantControllers.Retrieve)
	parentRouter.PUT("/variants/:id", variantControllers.Update)
	parentRouter.DELETE("/variants/:id", variantControllers.Delete)

	parentRouter.POST("/schemas", schemaControllers.Create)
	parentRouter.GET("/schemas", schemaControllers.List)
	parentRouter.GET("/schemas/:id", schemaControllers.Retrieve)
	parentRouter.PUT("/schemas/:id", schemaControllers.Update)
	parentRouter.DELETE("/schemas/:id", schemaControllers.Delete)

	parentRouter.POST("/variant_schemas", variantSchemaControllers.Create)
}
