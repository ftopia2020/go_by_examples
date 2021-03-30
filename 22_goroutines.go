package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    f("direct")				// main() 函数同步调用 f() 函数

	// 协程(goroutine) 是轻量级的执行线程
    go f("goroutine")		// 使用 go f(s) 在一个协程中调用这个函数。 这个新的 Go 协程将会 并发地 执行这个函数。

    go func(msg string) {	// 匿名函数启动协程
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)	// 停顿一秒
    fmt.Println("done")		// 打印 done
}