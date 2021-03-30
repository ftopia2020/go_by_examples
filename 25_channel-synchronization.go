package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)				// 要在协程中运行函数 worker，done 通道将被用于通知其他协程这个函数已经完成工作

    <-done						// 程序将一直阻塞，直至收到 worker 使用通道发送的通知
}