package main

import (
	"context"
	pb "create-order-saga/proto/payment"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *server) ProcessPayment(ctx context.Context, req *pb.ProcessPaymentRequest) (*pb.ProcessPaymentResponse, error) {
	// Logika untuk memproses pembayaran
	return &pb.ProcessPaymentResponse{Status: "SUCCESS"}, nil
}

func (s *server) RefundPayment(ctx context.Context, req *pb.RefundPaymentRequest) (*pb.RefundPaymentResponse, error) {
	// Logika untuk mengembalikan pembayaran
	return &pb.RefundPaymentResponse{Status: "REFUNDED"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, &server{})
	log.Println("Payment Service running on :50052")
	grpcServer.Serve(lis)
}
