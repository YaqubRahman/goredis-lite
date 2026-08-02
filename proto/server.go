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

func (s *server) Get(ctx context.Context, req *GetRequest) (*GetResponse, error){
	value := s.store[req.Key]
	return &GetResponse{Value: value}, nil
}

func (s *server) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error){
	delete(s.store, req.Key)
	return &DeleteResponse{Success: true}, nil
}