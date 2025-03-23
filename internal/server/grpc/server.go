package grpc

import (
	"context"
	"net"
	"os"

	"github.com/slem7451/anti_bruteforce/internal/server"
	"github.com/slem7451/anti_bruteforce/internal/server/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	*grpc.Server
}

func NewServer(app server.App) *Server {
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, &service{app: app})
	reflection.Register(grpcServer)

	return &Server{
		Server: grpcServer,
	}
}

func (s *Server) Start(_ context.Context) error {
	listener, err := net.Listen("tcp", os.Getenv("GRPC_ADDRESS"))
	if err != nil {
		return err
	}

	return s.Server.Serve(listener)
}

func (s *Server) Stop(_ context.Context) error {
	s.Server.GracefulStop()
	return nil
}
