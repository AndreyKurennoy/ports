package main

import (
	"ports/server/config"
	"ports/server/application"
	"ports/server/interfaces/grpc"
)

//TODO:Add tests, linters
func main() {
	configuration, err := config.Parse()
	if err != nil {
		panic(err)
	}
	dependencies := application.InitDependencies()

	server := grpc.NewServer(configuration.GrpcClient, dependencies)
	server.Start()
}
