package main

import (
	"fmt"
	"time"
) 

func main() {
    messages := make(chan string)
    signals := make(chan bool)


	// 使用带一个 default 子句的 select 来实现 非阻塞 的发送、接收
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    msg := "hi"
    select {
	case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }
	msgS := <-messages 
	fmt.Println(msgS)

	// go func(){ messages <- "aaa" }()
	// go func(){ signals <- true }()
	// 非阻塞的多路 select
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
	case <- time.After(1 * time.Second):
		fmt.Println("timeout 1")
    default:
        fmt.Println("no activity")
    }
}