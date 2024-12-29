### 普通编译
```bash
go build -work -x main.go
```
详细信息查看 build-dynamic.log，可以查看$WORK目录下的中间文件

```bash
$ ldd main
	linux-vdso.so.1 (0x00007fde8cfe6000)
	libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007fde8cfba000)
	libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fde8cde6000)
	/lib64/ld-linux-x86-64.so.2 (0x00007fde8cfe8000)
```
说明这种方式产生的二进制依赖glibc动态库

### 静态编译
```
go build -work -x -ldflags "-linkmode external -extldflags -static" main.go
```
详细信息查看 build-static.log

```bash
ldd main
	not a dynamic executable
```

上面构建命令可以省略 `-linkmode external`，因为如果用户层Go代码中使用了cgo代码，
那么Go链接器将会自动选择外部链接机制。

外部链接机制则是Go链接器将所有的.o都写入一个.o文件中，再将其交给外部链接器（比如gcc或者clang）
去做最终的链接处理。