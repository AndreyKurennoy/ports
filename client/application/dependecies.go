package application

import (
	"ports/client/infrastructure/persistence"
	"google.golang.org/grpc"
)

type Dependencies struct {
	PortService
}

func InitDependencies(grpc *grpc.ClientConn) Dependencies {
	d := Dependencies{}
	d.PortService = CreatePortService(persistence.NewPortRepository(grpc))

	return d
}
