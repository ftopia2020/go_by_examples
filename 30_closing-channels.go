package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)		// 用于标识关闭通道

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)	// 关闭一个通道，不能再向这个通道写入值。 该特性可以向通道的接收方传达工作已经完成 done！
    fmt.Println("sent all jobs")

    <-done		// 程序将一直阻塞，直至收到 done 标识
}