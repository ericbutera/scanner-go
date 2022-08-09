package grpc

import (
	"log"
	"net"

	pb "storage-api/protobuf"

	"golang.org/x/net/context"
	grpcServer "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received request: %v", in.Name)
	return &pb.HelloReply{Message: "SERVER REPLY: Hello " + in.Name}, nil
}

const Address string = ":50051"

func Serve() {
	lis, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("grpc listening: %v", Address)

	s := grpcServer.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Print("Registered Greeter")

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Print("gRPC reflection registered")
}
