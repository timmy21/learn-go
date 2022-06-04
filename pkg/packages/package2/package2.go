package package2

import "github.com/timmy21/learn-go/pkg/packages/package2/internal/package1"

func Avg(first int, remains ...int) int {
	sum := package1.Sum(append([]int{first}, remains...)...)
	return sum / (1 + len(remains))
}
