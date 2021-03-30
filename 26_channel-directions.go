package main

import "fmt"

// 通道作为函数参数时，可以指定这个通道是否为只读或只写，该特性可以提升程序的类型安全
func ping(pings chan<- string, msg string) {			// ping函数只允许 写入 channel
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {	// pong函数只允许 读取 channel
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}