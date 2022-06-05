package kvstore

import (
	"context"
	"sync"

	"github.com/pkg/errors"
)

var _ Backend = (*MemBackend)(nil)

type NotFoundError struct {
	key string
}

func (e *NotFoundError) Error() string {
	return "key: " + e.key + " not found"
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	var e *NotFoundError
	return errors.As(err, &e)
}

type MemBackend struct {
	mu   sync.Mutex
	data map[string][]byte
}

func NewMemBackend() *MemBackend {
	return &MemBackend{
		data: make(map[string][]byte),
	}
}

func (m *MemBackend) Set(_ context.Context, key string, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
	return nil
}

func (m *MemBackend) Get(_ context.Context, key string) ([]byte, error) {
	v, ok := m.data[key]
	if !ok {
		return nil, errors.WithStack(&NotFoundError{key: key})
	}
	return v, nil
}
