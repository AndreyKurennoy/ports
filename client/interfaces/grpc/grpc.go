package grpc

import (
	"google.golang.org/grpc"
	"ports/client/config"
)

func CreateFromConfiguration(config *config.Configuration) *grpc.ClientConn {
	client, err := grpc.Dial(config.GrpcClient.Address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return client
}
