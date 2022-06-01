// 在 Go 语言中，函数是一等公民类型，可以将使用当成值使用。这使得函数使用起来非常灵活。

package main

import (
	"fmt"
	"time"
)

// 支持高阶函数的用法，函数可以作为其他函数的参数，和返回值
func filter(nums []int, fn func(int) bool) []int {
	var result []int
	for _, num := range nums {
		if fn(num) {
			result = append(result, num)
		}
	}
	return result
}

// 可以自定义函数类型
type FilterFunc func(int) bool

func (f FilterFunc) Filter(nums []int) []int {
	return filter(nums, f)
}

type Client struct {
	host    string
	port    int
	timeout time.Duration

	// 函数可以作为结构体的字段
	backoff func(attempt int) time.Duration
}

// 这种传参方式类似于某些语言中的命名参数。在 Go 中称为：Functional options
func NewClient(options ...func(*Client)) *Client {
	// 可以设置默认值
	svr := &Client{
		host:    "localhost",
		port:    8000,
		timeout: 5 * time.Second,
		backoff: func(int) time.Duration {
			return time.Second
		},
	}
	for _, o := range options {
		o(svr)
	}
	return svr
}

func WithHost(host string) func(*Client) {
	return func(s *Client) {
		s.host = host
	}
}

func WithPort(port int) func(*Client) {
	return func(s *Client) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) func(*Client) {
	return func(s *Client) {
		s.timeout = timeout
	}
}

func WithBackoff(backoff func(attempt int) time.Duration) func(*Client) {
	return func(s *Client) {
		s.backoff = backoff
	}
}

func main() {
	fmt.Println(filter([]int{1, 2, 3, 4}, func(num int) bool {
		return num%2 == 0
	}))

	fn := FilterFunc(func(num int) bool {
		return num%2 == 0
	})
	fmt.Println(fn.Filter([]int{1, 2, 3, 4}))

	client := NewClient(
		WithHost("127.0.0.1"),
		WithTimeout(10*time.Second),
	)
	fmt.Printf("%+v\n", client)
}
