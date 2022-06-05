package kvstore

import (
	"context"

	"go.uber.org/zap"
)

type Backend interface {
	Set(ctx context.Context, key string, value []byte) error
	Get(ctx context.Context, key string) ([]byte, error)
}

type Service struct {
	lg      *zap.Logger
	backend Backend
}

func NewService(backend Backend, logger *zap.Logger) *Service {
	return &Service{backend: backend}
}
