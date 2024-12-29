### 普通编译
```bash
go build -work -x main.go
```
详细信息查看 build-dynamic.log

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