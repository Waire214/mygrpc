package main

import (
	"google.golang.org/grpc"
	"grpcUser/proto/pb"
	"grpcUser/server"
	"log"
	"net"
)

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	serve := server.NewServer(&pb.UserRequests{})

	pb.RegisterUserServiceServer(s, serve)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
