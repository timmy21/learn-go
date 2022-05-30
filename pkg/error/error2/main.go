package main

import (
	"errors"
	"fmt"
)

const maxUint64 = 1<<64 - 1

// error 本质就是一个值，可以进行相等检查，但不能携带任何上下文信息
// 这种方式也被称为“sentinel errors”。
// 在进行 io 数据读取时，经常会使用到的 io.EOF 就是这种 error 形式

// 实际开发中不太使用这种方式.
var (
	ErrSyntax = errors.New("invalid syntax")
	ErrRange  = errors.New("value out of range")
)

func ParseUint64(s string) (uint64, error) {
	cutoff := uint64(maxUint64/10 + 1)
	var n uint64
	for _, c := range []byte(s) {
		var d byte
		switch {
		case '0' <= c && c <= '9':
			d = c - '0'
		default:
			return 0, ErrSyntax
		}
		if n >= cutoff {
			return 0, ErrRange
		}

		n *= 10
		n1 := n + uint64(d)
		if n1 < n {
			return 0, ErrRange
		}
		n = n1
	}
	return n, nil
}

func main() {
	_, err := ParseUint64("18446744073709551616")

	// not good
	if err == ErrRange {
		fmt.Printf("%+v\n", err)
	}

	// good
	if errors.Is(err, ErrRange) {
		fmt.Printf("%+v\n", err)
	}
}
