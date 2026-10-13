package main

import (
	"bytes"
	"context"
	"net"
	"testing"
	"time"

	pb "asdf/example"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

type testServer struct {
	pb.UnimplementedExampleServiceServer
	fail     bool
	metadata metadata.MD
}

func (s *testServer) Example(req *pb.ExampleRequest, stream grpc.ServerStreamingServer[pb.ExampleReply]) error {
	s.metadata, _ = metadata.FromIncomingContext(stream.Context())
	if s.fail {
		return status.Error(codes.Internal, "test failure")
	}
	for _, reply := range []string{"first", "second"} {
		if err := stream.Send(&pb.ExampleReply{Reply: reply}); err != nil {
			return err
		}
	}
	return nil
}
func TestStreamCompletionAndMetadata(t *testing.T) {
	for _, fail := range []bool{false, true} {
		listener := bufconn.Listen(1024 * 1024)
		server := grpc.NewServer()
		service := &testServer{fail: fail}
		pb.RegisterExampleServiceServer(server, service)
		go server.Serve(listener)
		conn, err := grpc.NewClient("passthrough:///local-test", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return listener.Dial() }))
		if err != nil {
			t.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		var output bytes.Buffer
		err = callExample(ctx, pb.NewExampleServiceClient(conn), "test-model", "test-key", &output)
		cancel()
		conn.Close()
		server.Stop()
		listener.Close()
		if fail {
			if status.Code(err) != codes.Internal {
				t.Fatalf("expected server error, got %v", err)
			}
		} else {
			if err != nil {
				t.Fatal(err)
			}
			if output.String() != "first\nsecond\n" {
				t.Fatalf("unexpected output %q", output.String())
			}
		}
		if got := service.metadata.Get("baseten-authorization"); len(got) != 1 || got[0] != "Api-Key test-key" {
			t.Fatalf("auth %v", got)
		}
		if got := service.metadata.Get("baseten-model-id"); len(got) != 1 || got[0] != "model-test-model" {
			t.Fatalf("routing %v", got)
		}
	}
}
