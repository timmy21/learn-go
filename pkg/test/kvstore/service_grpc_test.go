package kvstore

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/testing/protocmp"

	"github.com/timmy21/learn-go/pkg/test/kvstore/kvstorepb"
)

func TestService_Get(t *testing.T) {
	backend := NewMemBackend()
	backend.Set(context.Background(), "k1", []byte("v1"))
	srv := NewService(backend, zap.L())
	t.Run("existed", func(t *testing.T) {
		item, err := srv.Get(context.Background(), &kvstorepb.Key{Name: "k1"})
		require.NoError(t, err)
		assert.Zero(t, cmp.Diff(&kvstorepb.Item{
			Key:   "k1",
			Value: []byte("v1"),
		}, item, protocmp.Transform()))
	})
	t.Run("miss", func(t *testing.T) {
		item, err := srv.Get(context.Background(), &kvstorepb.Key{Name: "k2"})
		require.Error(t, err)
		assert.Nil(t, item)
		st, ok := status.FromError(err)
		assert.True(t, ok)
		assert.Equal(t, codes.NotFound, st.Code())
	})
}

func TestService_Set(t *testing.T) {
	backend := NewMemBackend()
	srv := NewService(backend, zap.L())
	_, err := srv.Set(context.Background(), &kvstorepb.Item{Key: "k1", Value: []byte("v1")})
	require.NoError(t, err)

	v, err := backend.Get(context.Background(), "k1")
	require.NoError(t, err)
	assert.Equal(t, []byte("v1"), v)
}
