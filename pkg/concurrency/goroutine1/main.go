package main

import (
	"fmt"
	"sync"
	"time"
)

func Number(n int) <-chan int {
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
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("sum:", sum)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}
