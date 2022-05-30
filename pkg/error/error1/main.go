package main

import (
	"errors"
	"fmt"
)

const maxUint64 = 1<<64 - 1

// Go函数支持多值返回，如果需要返回错误，一般都是最后一个值表示error
func ParseUint64(s string) (uint64, error) {
	cutoff := uint64(maxUint64/10 + 1)
	var n uint64
	for _, c := range []byte(s) {
		var d byte
		switch {
		case '0' <= c && c <= '9':
			d = c - '0'
		default:
			return 0, errors.New("invalid syntax")
		}
		if n >= cutoff {
			return 0, errors.New("value out of range")
		}

		n *= 10
		n1 := n + uint64(d)
		if n1 < n {
			return 0, errors.New("value out of range")
		}
		n = n1
	}
	return n, nil
}

func main() {
	_, err := ParseUint64("18446744073709551616")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
