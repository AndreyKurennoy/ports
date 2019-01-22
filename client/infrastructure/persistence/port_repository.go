package persistence

import (
	"context"
	"ports/client/domain/repository"
	ports "ports/client/interfaces"
	"time"

	"google.golang.org/grpc"
)

// PortRepositoryImpl is abstraction over port model data storage.
type PortRepositoryImpl struct {
	Grpc ports.PortDomainClient
}

// NewPortRepository creates grpc repository for ports
func NewPortRepository(grpc *grpc.ClientConn) repository.PortRepository {
	portDomainClient := ports.NewPortDomainClient(grpc)
	return &PortRepositoryImpl{Grpc: portDomainClient}
}

// Get searches for port by id
func (r *PortRepositoryImpl) Get(id string) (*ports.Port, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	port, err := r.Grpc.GetPort(ctx, &ports.PortRequest{Id: id})
	if err != nil {
		//TODO: rewrite error for client
		return nil, err
	}

	return port, nil
}

// Save adds new ports to storage and updates existed
func (r *PortRepositoryImpl) Save(ports *ports.Ports) (*ports.Empty, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := r.Grpc.AddPorts(ctx, ports)
	if err != nil {
		//TODO: rewrite error for client
		return nil, err

	}

	return response, nil
}
