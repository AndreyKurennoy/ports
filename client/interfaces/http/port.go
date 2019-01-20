package http

import (
	"github.com/labstack/echo"
	"net/http"
	"ports/client/application"
)

type PortController struct {
	portService application.PortService
}

func (c *PortController) AddPorts(ctx echo.Context) error {

	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	_, err = c.portService.AddPorts(file)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (c *PortController) FindPort(ctx echo.Context) error {

	port, err := c.portService.FindPort(ctx.QueryParam("portId"))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, port)
}
