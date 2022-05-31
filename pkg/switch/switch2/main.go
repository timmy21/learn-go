package main

import (
	"fmt"
	"time"
)

// 首先执行 switch 初始化表达式。
// 然后从上到下逐个 case 条件语句，执行第一个结果为 true 的分支。
// 区别于某些语言需要显示调用 break，Go 默认会在 case 语句块的最后执行 break。
// 如果还需要执行下一个 case 语句块，可以调用 fallthrough。
func Greet(now time.Time) {
	switch hour := now.Hour(); {
	case hour < 12:
		fmt.Println("Good morning!")
		fallthrough
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

func main() {
	Greet(time.Now())

	fmt.Println("============")
	Greet(time.Date(2022, 6, 1, 10, 0, 0, 0, time.Local))
}
