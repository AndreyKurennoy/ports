package main

import (
	"ports/client/application"
	"ports/client/config"
	"ports/client/interfaces/grpc"
	"ports/client/interfaces/http"
)

//TODO:Add tests, linters, documentation
func main() {
	configuration, err := config.Parse()
	grpcClient := grpc.CreateFromConfiguration(configuration)
	dependencies := application.InitDependencies(grpcClient)

	if err != nil {
		panic(err)
	}

	server := http.NewServer(configuration.HTTPServer, dependencies)
	server.Start()
}
