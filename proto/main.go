package proto

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

	grpcServer := grpc.NewServer()

	// make will create a real, empty, ready to use map that I can write into
	// without make, the variable will exist but theres no map behind it
	kvStore := &server{store: make(map[string]string)}

	proto.RegisterKeyValueServiceServer(grpcServer, kvStore)

	if err := grpcServer.Serve(listener); err != nil{
		log.Fatal("Erro serving", err)
	}


}