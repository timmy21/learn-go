## 编译ctestlib
```bash
$ gcc -c -Wall -Werror -fpic -o ./ctestlib/test.o ./ctestlib/test.c
$ gcc -shared -o ./ctestlib/libtest.so ./ctestlib/test.o
```

## 参考链接
1. https://github.com/andreiavrammsd/cgo-examples/tree/master
1. [Linux 静态库和动态库](https://subingwen.cn/linux/library/)
