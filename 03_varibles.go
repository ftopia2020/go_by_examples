package main

import "fmt"

func main() {

    var a = "initial"		// var 变量关键字；= 赋值
    fmt.Println(a)

    var b, c int = 1, 2		// 显式声明变量类型，赋值多个
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int				// 声明变量，不同数据类型有其初始值
    fmt.Println(e)

    f := "apple"			// := 简短赋值方式
    fmt.Println(f)
}