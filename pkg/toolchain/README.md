## 常用命令

* go run: 用于执行 Go 程序，执行单个文件使用 `go run xxx.go`，如果 main 包存在多个文件，使用 `go run .` 执行。
* go build: 用于构建二进制可执行文件。
* go test: 用于执行单元测试和性能测试。
* go mod init: 在当前目录生成 go.mod 文件。
* go mod tidy: 通过分析当前项目的源码文件，添加缺失的依赖包，或者删除不再需要的依赖包。
* go get: 用于增加、更新 go.mod 需要的依赖包。实际开发中 `go mod init` 更常用。
* go install: 用于安装第三方 Go 程序到 `GOBIN` 目录下。

使用 "go help <command>" 查看命令的详细信息

## 进一步阅读

* [Using Go Modules](https://go.dev/blog/using-go-modules)
* [cmd/go](https://pkg.go.dev/cmd/go)