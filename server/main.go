package main

import (
	"context"
	"net"

	"github.com/yuki-toida/ca-micro-grpc/server/pb"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterTestServer(server, &Server{})
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}

type Server struct{}

func (s *Server) Get(_ context.Context, r *pb.Request) (*pb.Response, error) {
	return &pb.Response{Message: r.Message}, nil
}
