package http

import (
	"ports/client/application"
	"ports/client/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server is http server
type Server struct {
	echo   *echo.Echo
	config *config.HTTPServer
}

// NewServer created new server based on config and dependencies
func NewServer(config *config.HTTPServer, dep application.Dependencies) *Server {
	s := new(Server)
	s.config = config

	s.echo = echo.New()
	s.echo.Use(middleware.Logger())

	s.initRoutes(dep)

	return s
}

func (s *Server) initRoutes(dependencies application.Dependencies) {
	api := s.echo.Group("/api")

	v1 := api.Group("/v1")

	portController := PortController{portService: dependencies.PortService}
	v1.GET("/port", portController.FindPort)
	v1.POST("/ports", portController.AddPorts)
}

// Start server
func (s Server) Start() {
	s.echo.Logger.Fatal(s.echo.Start(s.config.Address))
}
