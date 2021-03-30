package main

import "fmt"

func main() {

    messages := make(chan string, 2)	// 创建字符串通道，缓存2个值

    messages <- "buffered"		// 通道已缓冲，接收值非并发
    messages <- "channel"

    fmt.Println(<-messages)		// 正常按顺序接收通道的2个值
    fmt.Println(<-messages)
}