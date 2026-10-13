package main

import (
	"context"
	"errors"
	"testing"

	pb "asdf/example"
	"google.golang.org/grpc/metadata"
)

type failingStream struct{ sent int }

func (s *failingStream) Send(*pb.ExampleReply) error { s.sent++; return errors.New("closed stream") }
func (*failingStream) SetHeader(metadata.MD) error   { return nil }
func (*failingStream) SendHeader(metadata.MD) error  { return nil }
func (*failingStream) SetTrailer(metadata.MD)        {}
func (*failingStream) Context() context.Context      { return context.Background() }
func (*failingStream) SendMsg(any) error             { return nil }
func (*failingStream) RecvMsg(any) error             { return nil }
func TestSendFailureStopsStream(t *testing.T) {
	stream := &failingStream{}
	err := (&server{}).Example(&pb.ExampleRequest{Message: "test"}, stream)
	if err == nil || err.Error() != "closed stream" {
		t.Fatalf("send failure suppressed: %v", err)
	}
	if stream.sent != 1 {
		t.Fatalf("sent %d replies after failure", stream.sent)
	}
}
