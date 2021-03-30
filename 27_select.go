package main

import (
    "fmt"
    "time"
)

// select（选择器）让你可同时等待多个通道操作。 将协程、通道和选择器结合，是 Go 的一个强大特性
func main() {

	bT := time.Now()            // 开始时间

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// select 将同时等待这两个值，各自完成打印
	for i := 0; i < 2; i++ {
		select {					// 类似于用于通信的switch语句
		case msg1 := <-c1:			// 每个 case 必须是一个通信操作，要么发送 要么接收
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	eT := time.Since(bT)			// 从开始到当前所消耗的时间
	fmt.Println("Run time: ", eT)	// 程序总共仅运行了2秒左右，因为以上协程是并发执行的
}