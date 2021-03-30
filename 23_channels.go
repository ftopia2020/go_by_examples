package main

import "fmt"

func main() {

    // channels 是连接多个协程的管道
    messages := make(chan string)       // 使用 make(chan val-type) 创建新通道

    go func() { messages <- "ping" }()  // 使用 channel <- 语法 发送新值到通道

    msg := <-messages                   // 使用 <- channel 语法从通道中 接收值
    
    fmt.Println(msg)    // 默认发送和接收都是阻塞的，此特性使我们不用任何其它同步操作， 就可在程序结尾等待消息 "ping"
}