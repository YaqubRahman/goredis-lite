package main

import (
	"kvstore/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Error listening", err)
	}

	// Creating the server object in memory
	// This is when I register my services onto it
	grpcServer := grpc.NewServer()

	// make will create a real, empty, ready to use map that I can write into
	// without make, the variable will exist but theres no map behind it
	kvStore := &server{store: make(map[string]entry)}

	proto.RegisterKeyValueServiceServer(grpcServer, kvStore)

	// This starts it running - running forever handling requests
	if err := grpcServer.Serve(listener); err != nil{
		log.Fatal("Erro serving", err)
	}


}