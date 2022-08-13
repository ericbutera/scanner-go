package main

import (
	"context"
	"log"
	pb "storage-client-grpc/protobuf"
	"time"

	grpc "google.golang.org/grpc"
)

func Client() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure()) // TODO add TLS
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
