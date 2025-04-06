package main

import (
	"context"
	pb "create-order-saga/proto/order"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedOrderServiceServer
}

func (s *server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// Logika untuk membuat order
	return &pb.CreateOrderResponse{Status: "PENDING"}, nil
}

func (s *server) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	// Logika untuk membatalkan order
	return &pb.CancelOrderResponse{Status: "CANCELLED"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &server{})
	log.Println("Order Service running on :50051")
	grpcServer.Serve(lis)
}
