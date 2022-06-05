package kvstore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"

	. "github.com/timmy21/learn-go/pkg/test/testutils"
)

// 对于外部数据库都可以采用类似的方式进行测试。
//   1. 创建一个测试数据库
//   2. 初始化相关的表结构（mongo不需要）
//   3. 使用测试数据库进行测试
//   4. 清理测试数据库
func TestMongoBackend(t *testing.T) {
	GetMongoClient(t).Tx(t, func(db *mongo.Database) {
		backend := NewMongoBackend(db)
		require.NoError(t, backend.Set(context.Background(), "k1", []byte("v1")))
		require.NoError(t, backend.Set(context.Background(), "k2", []byte("v2")))

		v1, err := backend.Get(context.Background(), "k1")
		require.NoError(t, err)
		assert.Equal(t, []byte("v1"), v1)

		_, err = backend.Get(context.Background(), "k3")
		require.True(t, IsNotFound(err))
	})
}
