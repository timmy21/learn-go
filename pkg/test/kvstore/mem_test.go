package kvstore

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemBackend_Get(t *testing.T) {
	backend := NewMemBackend()
	backend.Set(context.Background(), "k1", []byte("v1"))
	backend.Set(context.Background(), "k2", []byte("v2"))

	// 表驱动测试，减少重复代码，提升测试可读性
	tests := []struct {
		name     string
		key      string
		want     []byte
		checkErr func(t *testing.T, err error)
	}{
		{
			name: "existed",
			key:  "k1",
			want: []byte("v1"),
		},
		{
			name: "miss",
			key:  "k3",
			checkErr: func(t *testing.T, err error) {
				assert.True(t, IsNotFound(err))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := backend.Get(context.Background(), tt.key)
			if err != nil {
				tt.checkErr(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMemBackend_Set(t *testing.T) {
	backend := NewMemBackend()

	tests := []struct {
		key   string
		value []byte
	}{
		{
			key:   "k1",
			value: []byte("v1"),
		},
		{
			key:   "k2",
			value: []byte("v2"),
		},
		{
			key:   "k3",
			value: []byte("v3"),
		},
	}
	// Go 按次序逐个执行同一个包中的所有测试，使用 t.Parallel，本单元测试组下的子测试会并发执行。
	for _, tt := range tests {
		tt := tt
		t.Run(tt.key, func(t *testing.T) {
			t.Parallel()
			err := backend.Set(context.Background(), tt.key, tt.value)
			require.NoError(t, err)
		})
	}
}

// 性能测试
func BenchmarkMemBackend_Set(b *testing.B) {
	backend := NewMemBackend()
	for i := 0; i < b.N; i++ {
		backend.Set(context.Background(), "key-"+strconv.Itoa(i), []byte("value"+strconv.Itoa(i)))
	}
}
