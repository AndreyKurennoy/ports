package grpc

import (
	"fmt"
	"log"
	"net"
	"ports/server/application"
	"ports/server/config"
	ports "ports/server/interfaces"

	"google.golang.org/grpc"
)

// Server is gRPC server
type Server struct {
	config *config.GrpcServer
	gs     *grpc.Server
}

// NewServer created new server based on config and dependencies
func NewServer(config *config.GrpcServer, dependencies application.Dependencies) *Server {
	s := new(Server)
	s.config = config
	s.gs = grpc.NewServer()
	s.initRoutes(dependencies)

	return s
}

func (s *Server) initRoutes(dependencies application.Dependencies) {
	portController := PortController{portService: dependencies.PortService}
	ports.RegisterPortDomainServer(s.gs, &portController)
}

// Start server
func (s *Server) Start() {
	lis, err := net.Listen("tcp", s.config.Address)
	if err != nil {
		log.Fatalln("can't listen port", err)
	}
	fmt.Print("Server started!")
	log.Fatal(s.gs.Serve(lis))
}
