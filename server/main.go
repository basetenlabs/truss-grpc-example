package main

import (
	"fmt"
	"log"
	"net"

	pb "asdf/spacer"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type server struct {
	pb.UnimplementedYapperServiceServer
}

func (s *server) Yap(req *pb.YapRequest, stream grpc.ServerStreamingServer[pb.YapReply]) error {
	for i := 0; i < 100; i++ {
		stream.Send(&pb.YapReply{
			Reply: fmt.Sprintf("%d: %s", i, req.Message),
		})
	}
	return nil
}

var healthServer *health.Server
var healthStatus healthpb.HealthCheckResponse_ServingStatus

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}

	s := grpc.NewServer()

	healthServer = health.NewServer()

	healthStatus = healthpb.HealthCheckResponse_SERVING
	healthServer.SetServingStatus("yapper.YapperService", healthStatus)

	healthpb.RegisterHealthServer(s, healthServer)

	pb.RegisterYapperServiceServer(s, &server{})

	serviceInfo := s.GetServiceInfo()
	fmt.Println("Registered gRPC services:")
	for serviceName := range serviceInfo {
		fmt.Println("- ", serviceName)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
