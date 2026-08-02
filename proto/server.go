package proto

import "context"

type server struct {
	UnimplementedKeyValueServiceServer
	store map[string]string
}

func (s *server) Set(ctx context.Context, req *SetRequest) (*SetResponse, error) {
	s.store[req.Key] = req.Value
	return &SetResponse{Success: true}, nil
}
