package main

import (
	"context"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
	"net"
	"os"
	"server/pb"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	// The gRPC listener mounts the Go kit gRPC server we created.
	grpcListener, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}
	defer grpcListener.Close()

	server := grpc.NewServer()
	pb.RegisterAddressServer(server, &Server{})

	logger.Log("err", server.Serve(grpcListener))
}

type Server struct{}

func (s *Server) Get(_ context.Context, r *pb.Request) (*pb.Response, error) {
	address := r.Email + "@hacobu.jp"
	return &pb.Response{EmailAddress: address}, nil
}
