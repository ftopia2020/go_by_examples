package main

import "fmt"

func main() {

    // map 为 golang 内置映射数据类型，或称为hash或词典dict类型
    m := make(map[string]int)   // make构建空map，make(map[key-type]val-type)

    m["k1"] = 7                 // 插入或修改map元素
    m["k2"] = 13

    fmt.Println("map:", m)      // 直接打印 map 类型

    v1 := m["k1"]               // 获取 map 元素的值
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m)) // map 数量

    delete(m, "k2")             // 删除 map 元素
    fmt.Println("map:", m)

    _, prs := m["k2"]           // elem, ok = m[key] 通过双赋值检测某个键是否存在
    fmt.Println("prs:", prs)    // false - 不存在

    n := map[string]int{"foo": 1, "bar": 2} // 初始化即赋值
    fmt.Println("map:", n)
}