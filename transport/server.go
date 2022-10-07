package server

import (
	"net/http"

	"github.com/aviralbansal29/duis/config"
	"github.com/aviralbansal29/duis/log"
	httpRoutes "github.com/aviralbansal29/duis/transport/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RestServer contains global objects
type RestServer struct {
	mux    *echo.Echo
	server *http.Server
}

// Setup creates router instance
func (s *RestServer) Setup() {
	s.mux = echo.New()
	s.mux.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowHeaders: []string{
			echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "X-Total-Count", echo.HeaderAuthorization,
		},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	s.server = &http.Server{
		Addr:    config.GetEnv().GetString("server_address"),
		Handler: s.mux,
	}
	httpRoutes.SetupRouter(s.mux)
}

// Start starts the server
func (s *RestServer) Start() {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.GetLogger().Fatal(err)
	}
}

// Stop stops the server
func (RestServer) Stop() {
}
