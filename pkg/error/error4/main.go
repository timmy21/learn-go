// 虽然 Go1.13 对错误处理进行了增强，但由于缺少堆栈信息，实际开发中依然使用 github.com/pkg/errors
package main

import (
	stderrs "errors"
	"fmt"
	"io"
	"strconv"

	"github.com/pkg/errors"
)

const maxUint64 = 1<<64 - 1

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
			return 0, errors.WithStack(&SyntaxError{Num: s})
		}
		if n >= cutoff {
			return 0, errors.WithStack(&RangeError{Num: s})
		}

		n *= 10
		n1 := n + uint64(d)
		if n1 < n {
			return 0, errors.WithStack(&RangeError{Num: s})
		}
		n = n1
	}
	return n, nil
}

type PercentError struct {
	err error
}

func (e *PercentError) Error() string {
	return "invalid percent: " + e.err.Error()
}

func (e *PercentError) Unwrap() error {
	return e.err
}

func (e *PercentError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = io.WriteString(s, e.Error()+"\n")
			_, _ = fmt.Fprintf(s, "%+v", e.err)
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", e.Error())
	}
}

func ParsePercent(s string) (uint8, error) {
	v, err := ParseUint64(s)
	if err != nil {
		// 不要重复添加堆栈信息
		return 0, &PercentError{err: err}
	}
	if v > 100 {
		return 0, errors.WithStack(&PercentError{
			err: stderrs.New("percent should bettween 0 and 100"),
		})
	}
	return uint8(v), nil
}

func main() {
	_, err := ParseUint64("18446744073709551616")

	// 无法使用这种方式进行检查了
	if _, ok := err.(*RangeError); ok {
		fmt.Printf("%+v\n", err)
	}

	// good，基于错误行为进行检查
	if IsRangeError(err) {
		fmt.Printf("%+v\n", err)
	}

	fmt.Println("================")

	// 如果 *PercentError 不实现自定义Format，下面 Printf 不会显示堆栈信息
	_, err = ParsePercent("10.1")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
