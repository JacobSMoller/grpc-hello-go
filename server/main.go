// Package main implements a server for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/JacobSMoller/grpc-hello-go/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	log.Printf("Received: %v, %v", in.Name, in.Age)
	response := fmt.Sprintf("Hello %s you are %d years old", in.Name, in.Age)
	return &pb.HelloReply{Message: response}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
