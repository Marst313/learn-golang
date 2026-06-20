package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "simplegrpcserver/proto/gen"
)

type server struct {
	pb.UnimplementedCalculateServer
	pb.UnimplementedGreeterServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {

	sum := req.A + req.B
	log.Println(sum)

	return &pb.AddResponse{
		Sum: sum,
	}, nil

}

func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	message := "Hello " + req.Name
	log.Println(message)

	return &pb.HelloResponse{
		Message: message,
	}, nil
}

func main() {
	cert := "cert.pem"
	key := "key.pem"

	port := ":50051"
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("Failed to listen", err)
	}

	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatal("Failed to add TLS", err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterCalculateServer(grpcServer, &server{})
	pb.RegisterGreeterServer(grpcServer, &server{})

	log.Println("Server is running on port", port)
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatal("Failed to serve", err)
	}

}
