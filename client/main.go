// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/JacobSMoller/grpc-hello-go/proto"
	"google.golang.org/grpc"
)

const (
	address            = "localhost:50051"
	defaultName        = "world"
	defaultAge  uint32 = 19
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	age := defaultAge
	if len(os.Args) > 1 {
		name = os.Args[1]
		if intAge, err := strconv.ParseInt(os.Args[2], 10, 32); err == nil {
			age = intAge
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	request := &pb.HelloRequest{Name: name, Age: age}
	log.Println(request)
	r, err := c.SayHello(ctx, request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
