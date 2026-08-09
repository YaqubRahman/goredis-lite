package main

import (
	"context"
	"kvstore/proto"
	"sync"
	"time"
)

type entry struct {
	value string
	expiry time.Time
}

type server struct {
	proto.UnimplementedKeyValueServiceServer
	mu sync.RWMutex
	store map[string]entry
}

func (s *server) Set(ctx context.Context, req *proto.SetRequest) (*proto.SetResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[req.Key] = entry{req.Value, time.Now().Add(10 * time.Second)}
	return &proto.SetResponse{Success: true}, nil
}

func (s *server) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error){
	s.mu.RLock()
	defer s.mu.RUnlock()
	e, ok := s.store[req.Key]
	if !ok || time.Now().After(e.expiry){
		return &proto.GetResponse{Ok: false}, nil
	}
	return &proto.GetResponse{Value: e.value, Ok: true}, nil
}

func (s *server) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error){
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, req.Key)
	return &proto.DeleteResponse{Success: true}, nil
}