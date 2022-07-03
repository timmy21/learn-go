// 永远不要在不知道 goroutine 合适终止，或者不知道如何终止的情况下，启动一个 goroutine。否则很可能会导致 goroutine 泄露，内存泄露。
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Echo1 struct {
	stopCh chan struct{}
	exitCh chan struct{}
}

func NewEcho1() *Echo1 {
	return &Echo1{
		stopCh: make(chan struct{}),
		exitCh: make(chan struct{}),
	}
}

func (e *Echo1) Start() {
	go func() {
		defer close(e.exitCh)
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		var i int
		for {
			select {
			case <-ticker.C:
				fmt.Printf("echo1 %d\n", i)
				i++
			case <-e.stopCh:
				fmt.Println("echo1 stopped")
				return
			}
		}
	}()
}

func (e *Echo1) Stop() {
	// 确保 Stop 仅执行一次，多次 close 会导致 panic
	close(e.stopCh)
	<-e.exitCh
}

type Echo2 struct {
	ctx    context.Context
	cancel context.CancelFunc
	exitCh chan struct{}
}

func NewEcho2() *Echo2 {
	ctx, cancel := context.WithCancel(context.Background())
	return &Echo2{
		ctx:    ctx,
		cancel: cancel,
		exitCh: make(chan struct{}),
	}
}

func (e *Echo2) Start() {
	go func() {
		defer close(e.exitCh)
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		var i int
		for {
			select {
			case <-ticker.C:
				fmt.Printf("echo2 %d\n", i)
				i++
			case <-e.ctx.Done():
				fmt.Println("echo2 stopped")
				return
			}
		}
	}()
}

func (e *Echo2) Stop() {
	// Stop 可以执行多次
	e.cancel()
	<-e.exitCh
}

type Echo3 struct {
	stopCh chan struct{}
	exitCh chan struct{}
	once   sync.Once
}

func NewEcho3() *Echo3 {
	return &Echo3{
		stopCh: make(chan struct{}),
		exitCh: make(chan struct{}),
	}
}

func (e *Echo3) Start() {
	go func() {
		defer close(e.exitCh)
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		var i int
		for {
			select {
			case <-ticker.C:
				fmt.Printf("echo3 %d\n", i)
				i++
			case <-e.stopCh:
				fmt.Println("echo3 stopped")
				return
			}
		}
	}()
}

func (e *Echo3) Stop() {
	// Stop 可以执行多次
	e.once.Do(func() {
		close(e.stopCh)
	})
	<-e.exitCh
}

func main() {
	e1 := NewEcho1()
	e1.Start()

	e2 := NewEcho2()
	e2.Start()

	e3 := NewEcho3()
	e3.Start()

	time.Sleep(time.Second)
	e1.Stop()
	e2.Stop()
	e3.Stop()
}
