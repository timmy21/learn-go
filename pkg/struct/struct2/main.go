// Go 仅支持组合，不支持继承。为了让组合更加方便和有用，内嵌类型就是其中一个重要的功能
package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	City string
	Addr string
}

type Person struct {
	firstName string
	lastName  string
	address   Address
}

func (p Person) Name() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) SetAddress(addr Address) {
	p.address = addr
}

// Go 语言根本没有继承的概念，嵌入字段实现了类型之间的组合，它不是面向对象中的继承。
// 类型嵌入的主要目的是为了将被内嵌类型的功能扩展到内嵌它的结构体类型中。
type Employee struct {
	*Person
	Salary int
}

func inspectMethod(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Println(t, "methods:")
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("  method#%d: %s\n", i, t.Method(i).Name)
	}
}

func main() {
	e := Employee{
		Person: &Person{
			firstName: "tony",
			lastName:  "li",
			address: Address{
				City: "Shanghai",
				Addr: "xxxxxx",
			},
		},
		Salary: 10000,
	}
	// 在 Go 语言规范中 e.Name() 被称为：promoted method
	// 与继承的关键区别在于：调用 e.Name() 时，方法接收者是"e.Person"，而不是"e"
	fmt.Println(e.Name()) // 等价于 e.Person.Name()
	fmt.Println(e.Person.Name())

	// 等价于 e.Person.SetAddress(...)
	e.SetAddress(Address{City: "Beijing", Addr: "xxxxxx"})

	// 在 Go 语言规范中 e.address 被称为：promoted field
	fmt.Println(e.address)
	fmt.Println(e.Person.address)

	fmt.Println("===============")
	inspectMethod(e)
}
