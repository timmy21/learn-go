package package3

// 模块级私有的 internal 包，仅能被直接父包及其子包中的代码引用
// module 目录下的 internal 包，可以被本模块内部的包导入，但不可以被其他模块导入
import (
	// "github.com/timmy21/learn-go/pkg/packages/package1/internal/package1" // 无法通过编译
	"github.com/timmy21/learn-go/internal/package1"
	"github.com/timmy21/learn-go/pkg/packages/package2"
)

type Nums []int

func (n Nums) Sum() int {
	return package1.Sum(n...)
}

func (n Nums) Avg() int {
	if len(n) == 0 {
		return 0
	}
	return package2.Avg(n[0], n[1:]...)
}
