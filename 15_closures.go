package main

import "fmt"

func intSeq() func() int {	// 闭包只接受匿名函数 func(...)
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())	// 在函数外部访问函数内部变量成为可能
    fmt.Println(nextInt())	// 闭包函数内部变量离开其作用域后始终保持在内存中而不被销毁

    newInts := intSeq()
    fmt.Println(newInts())
}