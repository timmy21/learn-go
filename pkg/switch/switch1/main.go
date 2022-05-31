// 使用 switch 代替 if - else 语句，可以写出更简单、可读性更强的代码。
package main

import "fmt"

type State int

const (
	PENDING State = iota
	STARTING
	RUNNING
	DONE
	ERROR
)

// 从上到下逐个比较，执行第一个等于 switch 表达式的 case 分支。
func (s State) String() string {
	switch s {
	case PENDING:
		return "PENDING"
	case STARTING:
		return "STARTING"
	case RUNNING:
		return "RUNNING"
	case DONE:
		return "DONE"
	case ERROR:
		return "ERROR"
	default:
		return fmt.Sprintf("State(%d)", s)
	}
}

func main() {
	s := RUNNING
	fmt.Println(s)
}
