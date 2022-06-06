package kvstore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/timmy21/learn-go/pkg/test/kvstore/kvstorepb"
	"github.com/timmy21/learn-go/pkg/test/testutils"
)

func TestGrpcClient_Set(t *testing.T) {
	s, conn, err := testutils.NewGrpcServer(context.Background())
	require.NoError(t, err)

	backend := NewMemBackend()
	srv := NewService(backend, zap.L())
	kvstorepb.RegisterKVStoreServer(s, srv)
	s.Start()
	defer s.Stop()

	client, err := NewGrpcClient(conn)
	require.NoError(t, err)
	err = client.Set(context.Background(), "k1", []byte("v1"))
	require.NoError(t, err)

	v, err := backend.Get(context.Background(), "k1")
	require.NoError(t, err)
	assert.Equal(t, []byte("v1"), v)
}

func TestGrpcClient_Get(t *testing.T) {
	s, conn, err := testutils.NewGrpcServer(context.Background())
	require.NoError(t, err)

	backend := NewMemBackend()
	backend.Set(context.Background(), "k1", []byte("v1"))
	srv := NewService(backend, zap.L())
	kvstorepb.RegisterKVStoreServer(s, srv)
	s.Start()
	defer s.Stop()

	client, err := NewGrpcClient(conn)
	require.NoError(t, err)
	v, err := client.Get(context.Background(), "k1")
	require.NoError(t, err)
	assert.Equal(t, []byte("v1"), v)

	v, err = client.Get(context.Background(), "k2")
	require.True(t, IsNotFound(err))
	require.Nil(t, v)
}
