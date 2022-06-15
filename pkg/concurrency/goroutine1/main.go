// Go 调度器(CPM)本质是把大量的 goroutine 调度到少量的内核线程上去执行，并利用多核并行处理，实现并发。
// 可以使用 GOMAXPROCS 限定 P 的数量，默认为：runtime.NumCPU
// goroutine 初始栈大小为 2KB，并在不够时进行扩容。
// https://github.com/golang/go/blob/master/src/runtime/stack.go#L75
// https://pkg.go.dev/runtime/debug#SetMaxStack
package main

import (
	"fmt"
	"sync"
	"time"
)

func Number(n int) <-chan int {
	// 不带缓冲的 channel
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func Partition(in <-chan int, n int) []<-chan int {
	outs := make([]<-chan int, n)
	for i := 0; i < n; i++ {
		out := make(chan int)
		outs[i] = out
		go func(out chan<- int) {
			var sum int
			for num := range in {
				sum += num
			}
			out <- sum
			close(out)
		}(out)
	}
	return outs
}

func Merge(ins ...<-chan int) <-chan int {
	// 带缓冲的 channel
	out := make(chan int, 5)
	var wg sync.WaitGroup
	for _, in := range ins {
		wg.Add(1)
		go func(in <-chan int) {
			defer wg.Done()
			for num := range in {
				out <- num
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	nums := Number(100)
	parts := Partition(nums, 5)
	ch := Merge(parts...)

	done := make(chan struct{})
	var sum int
	go func() {
		for num := range ch {
			sum += num
		}
		// 多次 close 会导致 panic
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("sum:", sum)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}
