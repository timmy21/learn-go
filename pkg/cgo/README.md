## docker环境
```bash
$ docker run -it --rm -v .:/opt -w /opt golang:1.23-bullseye /bin/bash
```
在docker环境中测试 gcc 部分


## 进一步阅读
* [cmd/cgo](https://pkg.go.dev/cmd/cgo)
* [C? Go? Cgo!](https://go.dev/blog/cgo)
* [cgo is not Go](https://dave.cheney.net/tag/cgo)
* [Go 静态编译 和 CGO](https://www.rectcircle.cn/posts/go-static-compile-and-cgo/)
* [Go与C的桥梁：CGO入门剖析与实践](https://mp.weixin.qq.com/s/AMv5IVBPU2lAY_qUwskk4g)
* [CGO内存模型](https://chai2010.cn/advanced-go-programming-book/ch2-cgo/ch2-07-memory.html) 
* [chdb-go](https://github.com/chdb-io/chdb-go)
* [动态链接详解-Golang](https://www.rectcircle.cn/posts/linux-dylib-detail-5-lang-go/)