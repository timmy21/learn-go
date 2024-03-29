由于 Go1.18 加入了泛型，接口类型可以有两个用途：值类型和类型约束。在 Go1.18 之前所有的接口类型都是值类型，但加入泛型后，有些接口类型只能用做"类型约束"。本章节介绍的接口类型仅包括“值类型”部分，有关"类型约束"部分请查看 [泛型](../generic/README.md)

## 进一步阅读

* [Go Data Structures: Interfaces](https://research.swtch.com/interfaces)
* [Why is there no type inheritance?](https://go.dev/doc/faq#inheritance)
* [go-internals: 接口](https://github.com/go-internals-cn/go-internals/blob/master/chapter2_interfaces/README.md)
* [Go 和 interface 探究](https://xargin.com/go-and-interface/)