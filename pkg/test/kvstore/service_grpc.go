package kvstore

import (
	"context"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/timmy21/learn-go/pkg/test/kvstore/kvstorepb"
)

// 对比 grpc 和 http，可以很容易发现 gpc 实现更加简单，客户端集成更加方便。
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
	if item.Key == "" {
		st := status.New(codes.InvalidArgument, "item key is empty")
		ds, err := st.WithDetails(&errdetails.BadRequest_FieldViolation{
			Field:       "key",
			Description: "item key is empty",
		})
		if err != nil {
			return nil, st.Err()
		}
		return nil, ds.Err()
	}
	err := s.backend.Set(ctx, item.Key, item.Value)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
