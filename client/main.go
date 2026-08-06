package main

import (
	"context"
	"kvstore/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// local development without TLS, we pass "insecure" credentials
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error connecting", err)
	}

	defer conn.Close()


	client := proto.NewKeyValueServiceClient(conn)

	_, err = client.Set(context.Background(), &proto.SetRequest{Key: "ninja", Value: "cool"})
	if err != nil{
		log.Fatal("Error with set response", err)
	}

	get, err := client.Get(context.Background(), &proto.GetRequest{Key: "ninja"})
	if err != nil {
		log.Fatal("Error with get response", err)
	}


	log.Println("Got value:", get.Value, "found:", get.Ok)


}

