package main

import (
	"context"
	orderpb "create-order-saga/proto/order"
	paymentpb "create-order-saga/proto/payment"
	shippingpb "create-order-saga/proto/shipping"
	"log"

	"google.golang.org/grpc"
)

func main() {
	orderConn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	paymentConn, _ := grpc.Dial("localhost:50052", grpc.WithInsecure())
	shippingConn, _ := grpc.Dial("localhost:50053", grpc.WithInsecure())
	defer orderConn.Close()
	defer paymentConn.Close()
	defer shippingConn.Close()

	orderClient := orderpb.NewOrderServiceClient(orderConn)
	paymentClient := paymentpb.NewPaymentServiceClient(paymentConn)
	shippingClient := shippingpb.NewShippingServiceClient(shippingConn)

	// Langkah 1: Panggil Order Service
	orderResp, err := orderClient.CreateOrder(context.Background(), &orderpb.CreateOrderRequest{UserId: 1}) // Perbaikan di sini
	if err != nil {
		log.Fatalf("Error creating order: %v", err)
	}

	// Langkah 2: Panggil Payment Service
	paymentResp, err := paymentClient.ProcessPayment(context.Background(), &paymentpb.ProcessPaymentRequest{OrderId: orderResp.OrderId, Amount: 100.0}) // Menggunakan OrderId dari response
	if err != nil {
		log.Printf("Payment failed, refunding order...")
		orderClient.CancelOrder(context.Background(), &orderpb.CancelOrderRequest{OrderId: orderResp.OrderId}) // Menggunakan OrderId dari response
		return
	}

	// Langkah 3: Panggil Shipping Service
	shippingResp, err := shippingClient.StartShipping(context.Background(), &shippingpb.StartShippingRequest{OrderId: orderResp.OrderId}) // Menggunakan OrderId dari response
	if err != nil {
		log.Printf("Shipping failed, refunding payment...")
		paymentClient.RefundPayment(context.Background(), &paymentpb.RefundPaymentRequest{OrderId: orderResp.OrderId}) // Menggunakan OrderId dari response
		orderClient.CancelOrder(context.Background(), &orderpb.CancelOrderRequest{OrderId: orderResp.OrderId})         // Menggunakan OrderId dari response
		return
	}

	log.Printf("Order Status: %s, Payment Status: %s, Shipping Status: %s", orderResp.Status, paymentResp.Status, shippingResp.Status)
}
