package application

import (
	"ports/client/infrastructure/persistence"

	"google.golang.org/grpc"
)

// Dependencies stores all application dependencies.
type Dependencies struct {
	PortService
}

// InitDependencies initializes application dependencies.
func InitDependencies(grpc *grpc.ClientConn) Dependencies {
	d := Dependencies{}
	d.PortService = CreatePortService(persistence.NewPortRepository(grpc))

	return d
}
