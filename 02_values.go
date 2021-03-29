package main

import "fmt"

func main() {

    fmt.Println("go" + "lang")	// 字符串可使用 "+" 连接

    fmt.Println("1+1 =", 1+1)	// 整形 "+" 为运算符加号
    fmt.Println("7.0/3.0 =", 7.0/3.0)	// 浮点型

    fmt.Println(true && false)	// 布尔型逻辑运算符 与 && 
    fmt.Println(true || false)	// 或 ||
    fmt.Println(!true)			// 非 ！
}