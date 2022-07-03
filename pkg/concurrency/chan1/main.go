package main

// chan 底层数据结构：https://github.com/golang/go/blob/go1.18.3/src/runtime/chan.go#L33
/*
type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}

type waitq struct {
	first *sudog
	last  *sudog
}
*/

// 具体内部实现机制，直接查看源码：
// 1. https://github.com/golang/go/blob/go1.18.3/src/runtime/chan.go
// 2. https://github.com/golang/go/blob/go1.18.3/src/runtime/select.go
func unbuffered() {
	ch := make(chan bool)
	go func() {
		ch <- true
	}()
	<-ch
}

func buffered() {
	ch := make(chan bool, 1)
	ch <- true
	go func() {
		<-ch
	}()
	ch <- true
}

func main() {
	unbuffered()
	buffered()
}
