package main

import (
	"context"
	pb "create-order-saga/proto/shipping"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedShippingServiceServer
}

func (s *server) StartShipping(ctx context.Context, req *pb.StartShippingRequest) (*pb.StartShippingResponse, error) {
	// Logika untuk memulai pengiriman
	return &pb.StartShippingResponse{Status: "SHIPPED"}, nil
}

func (s *server) CancelShipping(ctx context.Context, req *pb.CancelShippingRequest) (*pb.CancelShippingResponse, error) {
	// Logika untuk membatalkan pengiriman
	return &pb.CancelShippingResponse{Status: "CANCELLED"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterShippingServiceServer(grpcServer, &server{})
	log.Println("Shipping Service running on :50053")
	grpcServer.Serve(lis)
}
