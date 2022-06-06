package kvstore

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/timmy21/learn-go/pkg/test/kvstore/kvstorepb"
)

func (s *Service) Get(ctx context.Context, key *kvstorepb.Key) (*kvstorepb.Item, error) {
	value, err := s.backend.Get(ctx, key.Name)
	switch {
	case IsNotFound(err):
		return nil, status.Error(codes.NotFound, err.Error())
	case err != nil:
		return nil, err
	default:
		return &kvstorepb.Item{Key: key.Name, Value: value}, nil
	}
}
func (s *Service) Set(ctx context.Context, item *kvstorepb.Item) (*emptypb.Empty, error) {
	err := s.backend.Set(ctx, item.Key, item.Value)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
