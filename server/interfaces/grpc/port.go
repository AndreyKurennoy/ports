package grpc

import (
	"ports/server/application"
	ports "ports/server/interfaces"

	"golang.org/x/net/context"
)

// PortController is controller
type PortController struct {
	portService application.PortService
}

// GetPort is http handler
func (c *PortController) GetPort(ctx context.Context, request *ports.PortRequest) (*ports.Port, error) {
	return c.portService.FindPort(request.Id)
}

// AddPorts is http handler
func (c *PortController) AddPorts(ctx context.Context, request *ports.Ports) (*ports.Empty, error) {
	return c.portService.AddPorts(request)
}
