package grpc

import (
	"ports/server/config"
	"log"
	"net"
	"google.golang.org/grpc"
	ports "ports/server/interfaces"
	"ports/server/application"
	"fmt"
)

type Server struct {
	config *config.GrpcServer
	gs     *grpc.Server
}

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

func (s *Server) Start() {
	lis, err := net.Listen("tcp", s.config.Address)
	if err != nil {
		log.Fatalln("can't listen port", err)
	}
	fmt.Print("Server started!")
	log.Fatal(s.gs.Serve(lis))
}
