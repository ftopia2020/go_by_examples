package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)		// 协程等待 2s
        c1 <- "result 1"
    }()

	// 超时 对于一个需要连接外部资源， 或者有耗时较长的操作的程序而言是很重要的
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):	// <-Time.After 等待超时 1s
        fmt.Println("timeout 1")		// 超时打印
    }

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)				// 未超时已接收，打印"result2"
    case <-time.After(3 * time.Second): // <-Time.After 等待超时 2s
        fmt.Println("timeout 2")
    }
}