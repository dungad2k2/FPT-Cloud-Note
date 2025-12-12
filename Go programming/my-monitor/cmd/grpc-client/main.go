package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "my-monitor/pkg/api/v1"
)

func main() {
	// 1. Connect to the server
	// We use "WithTransportCredentials(insecure)" because we don't have TLS/SSL set up locally.
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 2. Create the Client Stub
	client := pb.NewMonitorServiceClient(conn)

	// 3. Call the remote procedure
	// Notice we pass a Context! gRPC relies heavily on context for timeouts.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.CheckWebsite(ctx, &pb.CheckRequest{Url: "https://www.google.com"})
	if err != nil {
		log.Fatalf("could not check website: %v", err)
	}

	log.Printf("Result: %s is %s (Latency: %dms)", response.Url, response.Status, response.LatencyMs)
}