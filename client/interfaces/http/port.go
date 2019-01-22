package http

import (
	"net/http"
	"ports/client/application"

	"github.com/labstack/echo"
)

// PortController is controller
type PortController struct {
	portService application.PortService
}

// AddPorts is http handler
// POST /api/v1/ports
func (c *PortController) AddPorts(ctx echo.Context) error {

	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	_, err = c.portService.AddPorts(file)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, hi{
		"status": "ok",
	})
}

// FindPort is http handler
// GET /api/v1/port?portId={KEY}
func (c *PortController) FindPort(ctx echo.Context) error {

	port, err := c.portService.FindPort(ctx.QueryParam("portId"))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, port)
}
