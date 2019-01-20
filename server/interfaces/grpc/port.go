package grpc

import (
	"ports/server/application"
	ports "ports/server/interfaces"
	"golang.org/x/net/context"
)

type PortController struct {
	portService application.PortService
}

func (c *PortController) GetPort(ctx context.Context, request *ports.PortRequest) (*ports.Port, error) {
	return c.portService.FindPort(request.Id)
}

func (c *PortController) AddPorts(ctx context.Context, request *ports.Ports) (*ports.Empty, error) {
	return c.portService.AddPorts(request)
}
