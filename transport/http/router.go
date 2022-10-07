package httpRoutes

import (
	"fmt"
	"net/http"

	adminRoutes "github.com/aviralbansal29/duis/app/admin/routes"
	componentRoutes "github.com/aviralbansal29/duis/app/api/v1/component/routes"
	userRoutes "github.com/aviralbansal29/duis/app/api/v1/user/routes"
	"github.com/aviralbansal29/duis/config"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupRouter defines url endpoints
func SetupRouter(router *echo.Echo) {
	router.Pre(echoMiddleware.RemoveTrailingSlash())
	router.GET("/ping", func(c echo.Context) error {
		env := config.GetEnv()
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			env.GetString("postgres.host"),
			env.GetString("postgres.user"),
			env.GetString("postgres.password"),
			env.GetString("postgres.db"),
			env.GetString("postgres.port"),
		)
		_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return c.JSON(http.StatusBadGateway, map[string]string{"message": "Cannot connect to DB"})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Status OK"})
	})

	router.Use(echoMiddleware.RequestID())

	admin := router.Group("/admin")
	{
		admin.Use(echoMiddleware.Logger())
		adminRoutes.AddRoutes(admin)
	}

	apiRouter := router.Group("/api")
	{
		apiRouter.Use(echoMiddleware.Logger())
		v1 := apiRouter.Group("/v1")
		{
			userRoutes.AddRoutes(v1)
			componentRoutes.AddRoutes(v1)
		}
	}

	router.GET("/swagger/*", echoSwagger.WrapHandler)
}
