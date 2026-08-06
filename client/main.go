package main

import (
	"context"
	"kvstore/proto"
	"log"
	"time"

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

	// This adds a 5 second deadline if the calls take longer than 5 it cancels
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.Set(ctx, &proto.SetRequest{Key: "ninja", Value: "cool"})
	if err != nil{
		log.Fatal("Error with set response", err)
	}

	get, err := client.Get(ctx, &proto.GetRequest{Key: "ninja"})
	if err != nil {
		log.Fatal("Error with get response", err)
	}


	log.Println("Got value:", get.Value, "found:", get.Ok)


}

