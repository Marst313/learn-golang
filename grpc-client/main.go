package main

import (
	"context"
	"log"
	mainapipb "simplegrpcclient/proto/gen"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	cert := "cert.pem"

	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatal("Failed to load credentials", err)
	}

	// ! USE THIS IF NOT USING TLS
	// conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	// ! USE THIS IF USING TLS
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalln("Did not connect to server", err)
	}
	defer conn.Close()

	client := mainapipb.NewCalculateClient(conn)
	client2 := mainapipb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &mainapipb.AddRequest{
		A: 10,
		B: 20,
	}
	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalln("Could not add", err)
	}

	reqGreet := &mainapipb.HelloRequest{
		Name: "I Nyoman",
	}
	res2, err := client2.Greet(ctx, reqGreet)

	if err != nil {
		log.Fatalln("Could not greet", err)
	}

	log.Println("Sum:", res.Sum)
	log.Println("Greeting message", res2.Message)

	state := conn.GetState()
	log.Println("Connection state", state)

}
