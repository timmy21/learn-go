package kvstore

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

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
