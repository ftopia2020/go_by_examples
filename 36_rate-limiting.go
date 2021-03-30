package main

import (
    "fmt"
    "time"
)

func main() {

    bT := time.Now()            // 开始时间
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(200 * time.Millisecond)	// 每次间隔 200ms

    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }

    burstyLimiter := make(chan time.Time, 3)		// 最多允许3次并发

    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    go func() {
        for t := range time.Tick(500 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

    burstyRequests := make(chan int, 6)
    for i := 1; i <= 6; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("Burst request: ", req, time.Now())
    }
    eT := time.Since(bT)			// 从开始到当前所消耗的时间
	fmt.Println("Run time: ", eT)	// 程序总共仅运行了约1秒，协程是并发执行的
}