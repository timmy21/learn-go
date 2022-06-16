package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Go1.18 增加了“空接口”别名 any，推荐代码中使用 any 替换 interface{}
type any = interface{}

// 可用做值类型
type Reader interface {
	Read(p []byte) (n int, err error)
}

// 可用做值类型
type Writer interface {
	Write(p []byte) (n int, err error)
}

// 可用做值类型
type ReadWriter interface {
	Reader
	Writer
}

// 只能被用做类型约束
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// 只能被用做类型约束
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// 只能被用做类型约束
type Integer interface {
	Signed | Unsigned
}

// 只能被用做类型约束
type StringableSignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
	String() string
}

type MyInt int

func (mi MyInt) String() string {
	return fmt.Sprintf("MyInt(%d)", mi)
}

// 在Go中，接口实现是隐式的（非接口类型方法集包括接口方法集），不需要使用类似 implements 这样的关键字。这种方式也被称为：“duck typing”
func hello(v fmt.Stringer) string {
	return "Hello " + v.String() + "!"
}

type StatsReader struct {
	r io.Reader
	n int
}

func NewStatsReader(r io.Reader) *StatsReader {
	return &StatsReader{r: r}
}

// *StatsReader 隐式实现了 io.Reader 接口
func (r *StatsReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	r.n += n
	return
}

func main() {
	fmt.Println(hello(MyInt(1)))

	// 使用接口实现了“多态”特性
	r := NewStatsReader(strings.NewReader("10 11 12"))
	br := bufio.NewReader(r)
	for {
		v, err := br.ReadString(' ')
		fmt.Println(v)
		if err != nil {
			break
		}
	}
	fmt.Printf("total bytes: %d\n", r.n)
}
