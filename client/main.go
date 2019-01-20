package main

import (
	"ports/client/config"
	"ports/client/interfaces/grpc"
	"ports/client/application"
	"ports/client/interfaces/http"
)

func main() {
	configuration, err := config.Parse()
	grpcClient := grpc.CreateFromConfiguration(configuration)
	dependencies := application.InitDependencies(grpcClient)

	if err != nil {
		panic(err)
	}

	server := http.NewServer(configuration.HttpServer, dependencies)
	server.Start()
}
