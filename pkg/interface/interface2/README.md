
# 接口底层实现

## 非空接口

非空接口类型的内部定义如下：

```go
// 源码：https://github.com/golang/go/blob/go1.18.3/src/runtime/runtime2.go#L202
type iface struct {
	tab  *itab          // 类型、方法等信息
	data unsafe.Pointer // 动态值的原始数据的副本
}
```

```go
// 源码：https://github.com/golang/go/blob/go1.18.3/src/runtime/runtime2.go#L885
type itab struct {
	// 接口自己的静态类型
	inter *interfacetype
	// 动态值的类型
	_type *_type
	hash  uint32 // copy of _type.hash. Used for type switches.
	_     [4]byte
	// 实体类型和接口相关的方法，实体类型的其他方法并不会出现在这里
	// 存储的是第一个方法的函数指针，如果有更多的方法，在它之后的内存空间里继续存储。
	// 这些方法是按照函数名称的字典序进行排列的。
	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}
```

## 空接口

空接口类型的内部定义如下：
```go
// 源码：https://github.com/golang/go/blob/go1.18.3/src/runtime/runtime2.go#L207
type eface struct {
	_type *_type         // 动态值的类型
	data  unsafe.Pointer // 动态值的原始数据的副本
}
```
