package main

import (
	"log"
	"net"

	"example.com/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	service := NewEntityService()
	grpcServer := grpc.NewServer()

	pb.RegisterEntityServiceServer(grpcServer, service)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
