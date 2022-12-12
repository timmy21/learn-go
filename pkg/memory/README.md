## 内存释放

内存归还操作系统的详细流程，可以查看 `bgscavenge`：https://github.com/golang/go/blob/go1.18.3/src/runtime/mgcscavenge.go#L259

运行时调用 `sysUnused` 将内存归还给操作系统，在linux系统中调用 `madvise`，具体代码：https://github.com/golang/go/blob/go1.18.3/src/runtime/mem_linux.go#L111

调用 `madvise` 时，使用 `MADV_DONTNEED` 策略，它不如 `MADV_FREE` 快，但可以更快的让 RSS 值变小。

```golang
func parsedebugvars() {
	// defaults
	debug.cgocheck = 1
	debug.invalidptr = 1
	if GOOS == "linux" {
		// On Linux, MADV_FREE is faster than MADV_DONTNEED,
		// but doesn't affect many of the statistics that
		// MADV_DONTNEED does until the memory is actually
		// reclaimed. This generally leads to poor user
		// experience, like confusing stats in top and other
		// monitoring tools; and bad integration with
		// management systems that respond to memory usage.
		// Hence, default to MADV_DONTNEED.
		debug.madvdontneed = 1
	}
    ...
}
```

## 进一步阅读

* [The Go Memory Model](https://go.dev/ref/mem)
* [Memory Models - Russ Cox](https://research.swtch.com/mm)
* [Language Mechanics On Stacks And Pointers](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html)
* [Language Mechanics On Escape Analysis](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-escape-analysis.html)
* [Garbage Collection In Go : Part I - Semantics](https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html)
* [Garbage Collection In Go : Part II - GC Traces](https://www.ardanlabs.com/blog/2019/05/garbage-collection-in-go-part2-gctraces.html)
* [Garbage Collection In Go : Part III - GC Pacing](https://www.ardanlabs.com/blog/2019/07/garbage-collection-in-go-part3-gcpacing.html)
* [Visualizing memory management in Golang](https://deepu.tech/memory-management-in-golang/)
* [madvise](https://man7.org/linux/man-pages/man2/madvise.2.html)