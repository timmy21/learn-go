## 默认构建
```bash
$ go build main.go
$ ldd main
	linux-vdso.so.1 (0x00007f8c550a3000)
	libresolv.so.2 => /lib/x86_64-linux-gnu/libresolv.so.2 (0x00007f8c5507f000)
	libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007f8c5505d000)
	libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f8c54e89000)
	/lib64/ld-linux-x86-64.so.2 (0x00007f8c550a5000)
```

Go 应用由用户Go代码和Go标准库/运行时库组成，默认情况下，Go的运行时环境变量CGO_ENABLED=1，即默认开启cgo，
允许你在Go代码中调用C代码。Go的预编译标准库的.a文件也是在开启cgo的情况下编译出来的。

在 Go1.20 中不再预编译标准库的.a文件，详情查看：https://go.dev/doc/go1.20

在CGO_ENABLED=1时，编译标准库的net、os/user等几个标准库中的依赖cgo的包，Go链接器默认使用内部链接，
而无须启动外部链接器（比如gcc、clang等）。此时编译出来的最终二进制文件是动态链接，
这个时候即便传入`-extldflags -static`也是如此，因为根本没有用到外部链接器

```bash
$ go build -ldflags '-extldflags -static' main.go
$ ldd main
	linux-vdso.so.1 (0x00007fa79da94000)
	libresolv.so.2 => /lib/x86_64-linux-gnu/libresolv.so.2 (0x00007fa79da70000)
	libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007fa79da4e000)
	libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fa79d87a000)
	/lib64/ld-linux-x86-64.so.2 (0x00007fa79da96000)
```

外部链接机制则是Go链接器将所有的.o都写入一个.o文件中，再将其交给外部链接器（比如gcc或者clang）
去做最终的链接处理。如果此时传入`-extldflags -static`，那么gcc/clang将会做静态链接

```bash
$ go build -ldflags "-linkmode external -extldflags -static" main.go
# command-line-arguments
/usr/bin/ld: /tmp/go-link-3840791119/000004.o: in function `_cgo_04fbb8f65a5f_C2func_getaddrinfo':
/tmp/go-build/cgo-gcc-prolog:60: warning: Using 'getaddrinfo' in statically linked applications requires at runtime the shared libraries from the glibc version used for linking

$ ldd main
	not a dynamic executable
```

也可以使用CGO_ENABLED=0，使用纯Go版本的net包

```bash
$ CGO_ENABLED=0 go build main.go
$ ldd main
	not a dynamic executable
```