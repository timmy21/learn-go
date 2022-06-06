package kvstore

import (
	"context"

	"go.uber.org/zap"

	"github.com/timmy21/learn-go/pkg/test/kvstore/kvstorepb"
)

type Backend interface {
	Set(ctx context.Context, key string, value []byte) error
	Get(ctx context.Context, key string) ([]byte, error)
}

type Service struct {
	kvstorepb.UnimplementedKVStoreServer

	lg      *zap.Logger
	backend Backend
}

func NewService(backend Backend, logger *zap.Logger) *Service {
	return &Service{backend: backend}
}
