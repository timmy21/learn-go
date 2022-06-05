由于 Go1.18 加入了泛型，接口类型可以有两个用途：值类型和类型约束。在 Go1.18 之前所有的接口类型都是值类型，但加入泛型后，有些接口类型只能作为"类型约束"。本章节介绍的接口类型仅包括“值类型”部分，有关"类型约束"部分请查看 [泛型](../generic/README.md)

## 进一步阅读

* [Why is there no type inheritance?](https://go.dev/doc/faq#inheritance)