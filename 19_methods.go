package main

import "fmt"

type rect struct {
    width, height int
}

// 支持对结构体指针定义方法，接收类型 *rect
func (r *rect) area() int {
    return r.width * r.height
}

// 支持对结构体定义方法，接收类型 rect
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r	// Go 会自动处理值和指针之间的转换
	// 想要避免在调用方法时产生一个拷贝，或者想让方法可以修改接受结构体的值，建议使用指针来调用方法
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}