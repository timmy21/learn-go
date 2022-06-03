// 一个数组类型定义包括：特定长度和元素类型，长度不同或者元素类型不同的数组是不同类型。
// 计算机会为数组分配一块连续的内存来保存其中的元素。
package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

var randSource = rand.New(rand.NewSource(time.Now().UnixNano()))

type UUID [16]byte

func (u UUID) String() string {
	return hex.EncodeToString(u[:])
}

func NewID() UUID {
	var id [16]byte
	_, _ = randSource.Read(id[:])
	return id
}

// 数组变量包含整个数组，而不是引用或者第一个元素的指针
// 将数组作为函数参数，会导致整个数组被拷贝
func ReverseID(id UUID) UUID {
	for i := len(id)/2 - 1; i >= 0; i-- {
		j := len(id) - i - 1
		id[i], id[j] = id[j], id[i]
	}
	return id
}

func main() {
	id1 := NewID()
	fmt.Println(id1)

	id2 := ReverseID(id1)
	fmt.Println(id2)
	fmt.Println(id1)
}
