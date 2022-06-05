package main

import (
	"errors"
	"fmt"
	"strconv"
)

const maxUint64 = 1<<64 - 1

// 自定义错误类型，需要实现 error 接口
// https://github.com/golang/go/blob/go1.18.3/src/builtin/builtin.go#L270
type error interface {
	Error() string
}

type SyntaxError struct {
	Num string
}

func (e *SyntaxError) Error() string {
	return "parsing " + strconv.Quote(e.Num) + ": invalid syntax"
}

func IsSyntaxError(err error) bool {
	if err == nil {
		return false
	}
	var e *SyntaxError
	return errors.As(err, &e)
}

type RangeError struct {
	Num string
}

func (e *RangeError) Error() string {
	return "parsing " + strconv.Quote(e.Num) + ": value out of range"
}

func IsRangeError(err error) bool {
	if err == nil {
		return false
	}
	var e *RangeError
	return errors.As(err, &e)
}

func ParseUint64(s string) (uint64, error) {
	cutoff := uint64(maxUint64/10 + 1)
	var n uint64
	for _, c := range []byte(s) {
		var d byte
		switch {
		case '0' <= c && c <= '9':
			d = c - '0'
		default:
			return 0, &SyntaxError{Num: s}
		}
		if n >= cutoff {
			return 0, &RangeError{Num: s}
		}

		n *= 10
		n1 := n + uint64(d)
		if n1 < n {
			return 0, &RangeError{Num: s}
		}
		n = n1
	}
	return n, nil
}

func main() {
	_, err := ParseUint64("18446744073709551616")

	// not good
	if _, ok := err.(*RangeError); ok {
		fmt.Printf("%+v\n", err)
	}

	// good，基于错误行为进行检查
	if IsRangeError(err) {
		fmt.Printf("%+v\n", err)
	}
}
