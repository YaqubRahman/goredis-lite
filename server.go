package main

import (
	"context"
	"kvstore/proto"
	"sync"
)

type server struct {
	proto.UnimplementedKeyValueServiceServer
	mu sync.RWMutex
	store map[string]string
}

func (s *server) Set(ctx context.Context, req *proto.SetRequest) (*proto.SetResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[req.Key] = req.Value
	return &proto.SetResponse{Success: true}, nil
}

func (s *server) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error){
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.store[req.Key]
	return &proto.GetResponse{Value: value, Ok: ok}, nil
}

func (s *server) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error){
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, req.Key)
	return &proto.DeleteResponse{Success: true}, nil
}