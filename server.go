package main

import (
	"context"
	"kvstore/proto"
)

type server struct {
	proto.UnimplementedKeyValueServiceServer
	store map[string]string
}

func (s *server) Set(ctx context.Context, req *proto.SetRequest) (*proto.SetResponse, error) {
	s.store[req.Key] = req.Value
	return &proto.SetResponse{Success: true}, nil
}

func (s *server) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error){
	value, ok := s.store[req.Key]
	return &proto.GetResponse{Value: value, Ok: ok}, nil
}

func (s *server) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error){
	delete(s.store, req.Key)
	return &proto.DeleteResponse{Success: true}, nil
}