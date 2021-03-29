package main

import "fmt"

// 定义结构体（可理解为类class） type {struct_name} struct
type person struct {	
    name string
    age  int
}

func newPerson(name string) *person {	// 相当于构造函数
    p := person{name: name}
    p.age = 42
    return &p
}

func main() {

    fmt.Println(person{"Bob", 20})		// 创建结构体元素，支持打印

    fmt.Println(person{name: "Alice", age: 30})

    fmt.Println(person{name: "Fred"})	// 可省略部分初始值，默认对应数据类型的零值

    fmt.Println(&person{name: "Ann", age: 40}) // & 前缀生成一个结构体指针

    fmt.Println(newPerson("Jon"))		// 继续省略

    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)					// 获取结构体对应元素值，使用 .

    sp := &s							// 可以对结构体指针使用，指针会被自动解引用
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)
}