package main

import "fmt"

func main() {

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)                // 非空的通道也是可以关闭

    for elem := range queue {   // for ... range 遍历通道取值
        fmt.Println(elem)
    }

    // queue <- "xxx"           // panic: send on closed channel
}