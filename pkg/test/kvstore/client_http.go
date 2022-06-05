package kvstore

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/pkg/errors"
)

type HttpClient struct {
	client *http.Client
	target *url.URL
}

func NewHttpClient(endpoint string) (*HttpClient, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &HttpClient{
		client: &http.Client{},
		target: u,
	}, nil
}

func (c *HttpClient) Set(ctx context.Context, key string, value []byte) error {
	b, err := json.Marshal(struct {
		Key   string `json:"key"`
		Value Value  `json:"value"`
	}{
		Key:   key,
		Value: value,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	u := *c.target
	u.Path = path.Join(u.Path, "/api/set")
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, u.String(), bytes.NewBuffer(b))
	if err != nil {
		return errors.WithStack(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	var e HTTPError
	if err := json.Unmarshal(body, &e); err != nil {
		return errors.Errorf("response status is %d, %v", resp.StatusCode, string(body))
	}
	return errors.New(e.Err)
}

func (c *HttpClient) Get(ctx context.Context, key string) ([]byte, error) {
	u := *c.target
	u.Path = path.Join(u.Path, "/api/get")
	u.RawQuery = url.Values{"key": []string{key}}.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var result struct {
			Value Value `json:"value"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, errors.WithStack(err)
		}
		return result.Value, nil
	case http.StatusNotFound:
		return nil, errors.WithStack(&NotFoundError{key: key})
	default:
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		var e HTTPError
		if err := json.Unmarshal(body, &e); err != nil {
			return nil, errors.Errorf("response status is %d, %v", resp.StatusCode, string(body))
		}
		return nil, errors.New(e.Err)
	}
}
