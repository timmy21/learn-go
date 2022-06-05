package kvstore

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"go.uber.org/zap"
)

var decoder = schema.NewDecoder()

func init() {
	decoder.SetAliasTag("json")
}

type Value []byte

func (v Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(base64.StdEncoding.EncodeToString(v))
}

func (v *Value) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	val, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	*v = val
	return nil
}

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

func (s *Service) HttpGet(w http.ResponseWriter, r *http.Request) {
	hp := NewHelper(w, r, s.lg)
	var params struct {
		Key string `json:"key"`
	}
	if !hp.DecodeQuery(&params) {
		return
	}
	value, err := s.backend.Get(r.Context(), params.Key)
	switch {
	case IsNotFound(err):
		hp.NotFound(err)
	case err != nil:
		hp.Abort(err)
	default:
		hp.JSON(http.StatusOK, struct {
			Value Value `json:"value"`
		}{
			Value: value,
		})
	}
}

func (s *Service) HttpSet(w http.ResponseWriter, r *http.Request) {
	hp := NewHelper(w, r, s.lg)
	var params struct {
		Key   string `json:"key"`
		Value Value  `json:"value"`
	}
	if !hp.DecodeJSON(&params) {
		return
	}
	err := s.backend.Set(r.Context(), params.Key, params.Value)
	if err != nil {
		hp.Abort(err)
		return
	}
	hp.NoContent()
}

type HTTPError struct {
	Err string `json:"error"`
}

type Helper struct {
	W  http.ResponseWriter
	R  *http.Request
	lg *zap.Logger
}

func NewHelper(w http.ResponseWriter, req *http.Request, logger *zap.Logger) *Helper {
	return &Helper{
		W:  w,
		R:  req,
		lg: logger,
	}
}

func (hp *Helper) BadRequest(err error) {
	hp.Fail(http.StatusBadRequest, err)
}

func (hp *Helper) Abort(err error) {
	hp.lg.Error("abort", zap.Error(err))
	hp.Fail(http.StatusInternalServerError, err)
}

func (hp *Helper) Fail(code int, err error) {
	hp.W.Header().Set("Content-Type", "application/json; charset=utf-8")

	data, err := json.Marshal(HTTPError{
		err.Error(),
	})
	if err != nil {
		panic(err)
	}

	hp.W.WriteHeader(code)
	_, err = hp.W.Write(data)
	if err != nil {
		hp.lg.Error("write fail", zap.Error(err))
	}
}

func (hp *Helper) JSON(code int, v any) {
	hp.W.Header().Set("Content-Type", "application/json; charset=utf-8")

	data, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		hp.Abort(err)
		return
	}

	hp.W.WriteHeader(code)
	_, err = hp.W.Write(data)
	if err != nil {
		hp.lg.Error("write fail", zap.Error(err))
	}
}

func (hp *Helper) NotFound(err error) {
	hp.Fail(http.StatusNotFound, err)
}

func (hp *Helper) NoContent() {
	hp.W.WriteHeader(http.StatusNoContent)
	_, err := hp.W.Write(nil)
	if err != nil {
		hp.lg.Error("write fail", zap.Error(err))
	}
}

func (hp *Helper) DecodeQuery(v any) bool {
	err := decoder.Decode(v, hp.R.URL.Query())
	if err != nil {
		hp.BadRequest(err)
		hp.lg.Debug("decode query params", zap.Error(err))
		return false
	}
	return true
}

func (hp *Helper) DecodeJSON(v any) bool {
	err := json.NewDecoder(hp.R.Body).Decode(v)
	if err != nil {
		hp.BadRequest(err)
		hp.lg.Debug("parse json", zap.Error(err))
		return false
	}
	return true
}
