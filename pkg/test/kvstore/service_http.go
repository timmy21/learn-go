package kvstore

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/schema"
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

func NewMux(srv *Service) *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.RealIP,
	)
	r.Mount("/debug", middleware.Profiler())
	r.Route("/api", func(r chi.Router) {
		r.Get("/get", srv.HttpGet)
		r.Put("/set", srv.HttpSet)
	})
	return r
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
