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

func GetName(p Person) string {
	return p.firstName + " " + p.lastName
}

// 等价于上面的 GetName，调用 Name 方法时 Person 对象会进行拷贝。
// “方法接收者”名称通常使用1到2个字符，不要使用 this, self 这类命名。并且所有方法中使用同一个名字（注：可以不同，但不推荐）。
func (p Person) Name() string {
	return p.firstName + " " + p.lastName
}

func SetAddress(p *Person, addr Address) {
	p.address = addr
}

// 等价于上面的 SetAddress
func (p *Person) SetAddress(addr Address) {
	p.address = addr
}

// 值方法 vs 指针方法，在开发过程中可能会时常纠结使用哪一种，强烈建议阅读:
// https://github.com/golang/go/wiki/CodeReviewComments#receiver-type

func inspectMethod(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Println(t, "methods:")
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("  method#%d: %s\n", i, t.Method(i).Name)
	}
}

func main() {
	p := Person{
		firstName: "tony",
		lastName:  "li",
		address: Address{
			City: "Shanghai",
			Addr: "xxxxxx",
		},
	}
	pt := &p
	fmt.Println(p.Name())
	fmt.Println(pt.Name())      // 等价于: (*ptr).Name()
	fmt.Println(Person.Name(p)) // 不会使用这种方式

	newAddr := Address{City: "Beijing", Addr: "xxxxxx"}
	p.SetAddress(newAddr) // 等价于: (&p).SetAddress(...)
	pt.SetAddress(newAddr)
	// (*Person).SetAddress(p, newAddr) // 无法编译
	(*Person).SetAddress(pt, newAddr)
	fmt.Println(p.address)
	fmt.Println(pt.address)

	fmt.Println("================")
	// 之所以可以调用 p.SetAddress() ，是因为 p 是可寻址的。编译器会自动处理为：(&p).SetAddress(...)
	// 通过指针总是可以调用“值方法”的，因为通过指针总是可以获取到值对象。
	inspectMethod(p)
	inspectMethod(&p)
}
