// 使用 xxx_test 包，用于测试真实用户如何使用这个包，验证对外接口是否合理。
package kvstore_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/timmy21/learn-go/pkg/test/kvstore"
)

// 当依赖外部 http 服务时，通过下面方式模拟一个 http 测试服务。
// httptest.NewServer() 返回一个测试服务器，监听本地回环地址上的一个可用随机端口。
func newHttpTestServer() *httptest.Server {
	r := chi.NewRouter()
	r.Get("/api/get", func(w http.ResponseWriter, r *http.Request) {
		hp := kvstore.NewHelper(w, r, zap.L())
		var params struct {
			Key string `json:"key"`
		}
		if !hp.DecodeQuery(&params) {
			return
		}
		switch params.Key {
		case "k1":
			hp.JSON(http.StatusOK, struct {
				Value kvstore.Value `json:"value"`
			}{
				Value: kvstore.Value("v1"),
			})
		case "k2":
			hp.NotFound(errors.New("xxx"))
		default:
			hp.Abort(errors.New("xxx"))
		}
	})
	r.Put("/api/set", func(w http.ResponseWriter, r *http.Request) {
		hp := kvstore.NewHelper(w, r, zap.L())
		hp.NoContent()
	})
	return httptest.NewServer(r)
}

func TestHttpClient_Set(t *testing.T) {
	server := newHttpTestServer()
	defer server.Close()

	client, err := kvstore.NewHttpClient(server.URL)
	require.NoError(t, err)
	err = client.Set(context.Background(), "k1", []byte("v1"))
	require.NoError(t, err)
}

func TestHttpClient_Get(t *testing.T) {
	server := newHttpTestServer()
	defer server.Close()

	client, err := kvstore.NewHttpClient(server.URL)
	require.NoError(t, err)

	t.Run("ok", func(t *testing.T) {
		v, err := client.Get(context.Background(), "k1")
		require.NoError(t, err)
		assert.Equal(t, []byte("v1"), v)
	})
	t.Run("miss", func(t *testing.T) {
		v, err := client.Get(context.Background(), "k2")
		require.True(t, kvstore.IsNotFound(err))
		assert.Nil(t, v)
	})
	t.Run("error", func(t *testing.T) {
		v, err := client.Get(context.Background(), "k3")
		require.False(t, kvstore.IsNotFound(err))
		require.Error(t, err)
		assert.Nil(t, v)
	})
}
