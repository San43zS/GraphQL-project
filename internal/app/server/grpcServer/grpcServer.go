package grpcServer

import (
	"GraphQL-project/internal/config"
	"GraphQL-project/proto/api/generate/desc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	service desc.UserServiceServer
	config  *config.Config
}

func New(service desc.UserServiceServer, config *config.Config) *server {
	return &server{
		service: service,
		config:  config,
	}
}

func (s *server) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.GrpcPort))
	if err != nil {
		log.Fatalln("Failed to listen grpc server: ", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	desc.RegisterUserServiceServer(server, s.service)
	log.Printf("serving gRPC on http://localhost:%d\n", s.config.GrpcPort)

	errCh := make(chan error)
	go func() {
		defer close(errCh)
		errCh <- server.Serve(lis)
	}()
	select {
	case err = <-errCh:
		return err
	case <-ctx.Done():
		server.GracefulStop()
	}
	return nil
}
