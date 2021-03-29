package main

import (
    "fmt"
    "math"
)

type geometry interface {	// 方法签名的集合：接口(Interfaces)
    area() float64			// area()  面积方法
    perim() float64			// perim() 周长方法
}

type rect struct {			// 定义长方形rect结构体
    width, height float64
}
type circle struct {		// 定义圆形circle结构体
    radius float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {	// 测量方法，接收接口geometry
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}
    measure(r)
    measure(c)
}