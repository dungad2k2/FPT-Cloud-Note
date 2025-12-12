package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"
	
	// Import the generated code
	pb "my-monitor/pkg/api/v1"
)

// server controls the RPC service.
// It inherits from the Unimplemented struct for forward compatibility.
type server struct {
	pb.UnimplementedMonitorServiceServer
}

// CheckWebsite implements the logic defined in the .proto file
func (s *server) CheckWebsite(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	log.Printf("Received request for: %s", req.Url)

	// Simulate a check (or call your internal/checker logic here)
	start := time.Now()
	resp, err := http.Get(req.Url)
	latency := time.Since(start).Milliseconds()

	status := "DOWN"
	if err == nil {
		status = resp.Status
		resp.Body.Close()
	}

	// Return the Protobuf response
	return &pb.CheckResponse{
		Url:       req.Url,
		Status:    status,
		LatencyMs: int32(latency),
	}, nil
}

func main() {
	// 1. Listen on a TCP port
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 2. Create the gRPC Server
	grpcServer := grpc.NewServer()

	// 3. Register our service implementation
	pb.RegisterMonitorServiceServer(grpcServer, &server{})

	log.Println("🚀 gRPC Server listening on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}