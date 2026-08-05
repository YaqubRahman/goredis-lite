package main

import (
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

}

