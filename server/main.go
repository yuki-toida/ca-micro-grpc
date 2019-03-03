package main

import (
	"context"
	"net"

	"github.com/yuki-toida/grpc-clean/server/pb"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterTestServer(server, &Server{})
	server.Serve(listener)
}

type Server struct{}

func (s *Server) Get(_ context.Context, r *pb.Request) (*pb.Response, error) {
	return &pb.Response{Message: r.Message}, nil
}
