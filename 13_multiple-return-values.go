package main

import "fmt"

func vals() (int, int) {	// func 函数可接受多个值返回
    return 3, 7
}

func main() {

    a, b := vals()			// 多赋值 接收多个返回值
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals()			// 空白标识符忽略部分返回
    fmt.Println(c)
}