package main

import (
	"fmt"
	"net/netip"

	"github.com/pkg/errors"
)

var local netip.Addr

// 通常在 init 函数中进行初始化时，通常参数是受控的，不会出现错误，此时可以直接 panic
func init() {
	var err error
	local, err = netip.ParseAddr("127.0.0.1")
	if err != nil {
		panic(err)
	}
}

type TagMaker struct {
	Keys []string
}

// 在调用 API 时，如果确定参数不会导致错误，可以使用 panic 显示标识代码不可达，或者简化错误处理。
func (t TagMaker) New(vals ...string) map[string]string {
	if len(vals) != len(t.Keys) {
		panic(fmt.Sprintf("require %d value, but got %d", len(t.Keys), len(vals)))
	}
	m := make(map[string]string, len(t.Keys))
	for i := range vals {
		m[t.Keys[i]] = vals[i]
	}
	return m
}

// 在函数/库的内部在特定场景，可以谨慎的选择使用 panic，但对外 API 使用 error 返回值
func index(data any, idx any) (val any, err error) {
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		switch v := r.(type) {
		case error:
			err = v
		case string:
			err = errors.New(v)
		default:
			err = errors.Errorf("%v", v)
		}
	}()
	switch data := data.(type) {
	case []string:
		val = data[idx.(int)]
	case map[string]string:
		val = data[idx.(string)]
	default:
		err = errors.Errorf("unsupport data type: %T", data)
	}
	return
}

func main() {
	m := TagMaker{Keys: []string{"k1", "k2"}}
	fmt.Println(m.New("v1", "v2"))

	_, err := index([]string{"v1", "v2"}, "k1")
	if err != nil {
		fmt.Println(err)
	}
}
