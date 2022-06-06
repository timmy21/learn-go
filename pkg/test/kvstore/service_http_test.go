package kvstore

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// 服务提供方使用下面方式编写 http 测试
// 测试一个服务处理器，需要提供 *http.Request 对象和 http.ResponseWriter 接口。
//   1. httptest.NewRequest() 用于生成一个 *http.Request 对象
//   2. httptest.NewRecorder() 返回一个实现了 http.ResponseWriter 接口的 *httptest.ResponseRecorder 对象，用于记录处理器对 http.ResponseWriter 的修改
func TestService_HttpGet(t *testing.T) {
	backend := NewMemBackend()
	backend.Set(context.Background(), "k1", []byte("v1"))
	srv := NewService(backend, zap.L())
	t.Run("existed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.URL.RawQuery = url.Values{
			"key": []string{"k1"},
		}.Encode()

		w := httptest.NewRecorder()
		srv.HttpGet(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		var result struct {
			Value Value `json:"value"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &result)
		require.NoError(t, err)
		assert.Equal(t, Value("v1"), result.Value)
	})
	t.Run("miss", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.URL.RawQuery = url.Values{
			"key": []string{"k2"},
		}.Encode()

		w := httptest.NewRecorder()
		srv.HttpGet(w, req)

		require.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestService_HttpSet(t *testing.T) {
	backend := NewMemBackend()
	srv := NewService(backend, zap.L())

	body, err := json.Marshal(struct {
		Key   string `json:"key"`
		Value Value  `json:"value"`
	}{
		Key:   "k1",
		Value: Value("v1"),
	})
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(body))

	w := httptest.NewRecorder()
	srv.HttpSet(w, req)
	require.Equal(t, http.StatusNoContent, w.Code)

	val, err := backend.Get(context.Background(), "k1")
	require.NoError(t, err)
	assert.Equal(t, []byte("v1"), val)
}
