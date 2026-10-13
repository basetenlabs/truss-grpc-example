package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "asdf/example"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func callExample(ctx context.Context, client pb.ExampleServiceClient, modelID, apiKey string, output io.Writer) error {
	ctx = metadata.AppendToOutgoingContext(ctx,
		"baseten-authorization", "Api-Key "+apiKey,
		"baseten-model-id", "model-"+modelID,
	)
	stream, err := client.Example(ctx, &pb.ExampleRequest{Message: "😊"})
	if err != nil {
		return err
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if _, err := fmt.Fprintln(output, message.Reply); err != nil {
			return err
		}
	}
}

func main() {
	modelID, apiKey := os.Getenv("BASETEN_MODEL_ID"), os.Getenv("BASETEN_API_KEY")
	if modelID == "" || apiKey == "" {
		log.Fatal("Set BASETEN_MODEL_ID and BASETEN_API_KEY before running the client")
	}
	conn, err := grpc.NewClient("model-"+modelID+".grpc.api.baseten.co:443",
		grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := callExample(ctx, pb.NewExampleServiceClient(conn), modelID, apiKey, os.Stdout); err != nil {
		log.Fatal(err)
	}
}
