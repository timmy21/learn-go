package kvstore

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/timmy21/learn-go/pkg/test/kvstore/kvstorepb"
)

// GrpcClient 仅为了演示客户端使用，真实使用不需要这样包一层，直接使用 kvstorepb.KVStoreClient 即可
type GrpcClient struct {
	client kvstorepb.KVStoreClient
}

func NewGrpcClient(cc grpc.ClientConnInterface) (*GrpcClient, error) {
	return &GrpcClient{
		client: kvstorepb.NewKVStoreClient(cc),
	}, nil
}

func (c *GrpcClient) Set(ctx context.Context, key string, value []byte) error {
	_, err := c.client.Set(ctx, &kvstorepb.Item{Key: key, Value: value})
	st := status.Convert(err)
	if st.Code() == codes.InvalidArgument {
		for _, d := range st.Details() {
			switch info := d.(type) {
			case *errdetails.BadRequest_FieldViolation:
				return errors.Errorf("invalid field(%s): %s", info.Field, info.Description)
			}
		}
	}
	return errors.WithStack(err)
}

func (c *GrpcClient) Get(ctx context.Context, key string) ([]byte, error) {
	item, err := c.client.Get(ctx, &kvstorepb.Key{Name: key})
	if err == nil {
		return item.Value, nil
	}
	st := status.Convert(err)
	if st.Code() == codes.NotFound {
		return nil, errors.WithStack(&NotFoundError{key: key})
	}
	return nil, errors.WithStack(err)
}
