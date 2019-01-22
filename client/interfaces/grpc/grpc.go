package grpc

import (
	"ports/client/config"

	"google.golang.org/grpc"
)

// CreateFromConfiguration creates gRPC client connection using Configuration
func CreateFromConfiguration(config *config.Configuration) *grpc.ClientConn {
	client, err := grpc.Dial(config.GrpcClient.Address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return client
}
