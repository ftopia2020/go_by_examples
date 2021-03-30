package main

import (
    "fmt"
    "time"
)

func main() {

    ticker := time.NewTicker(500 * time.Millisecond)	// 500ms 打点
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:					// ticker 停止，跳出循环（通过读取 done 值，true 即满足条件）
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)	// ticker 打点
            }
        }
    }()

    time.Sleep(1600 * time.Millisecond)		// 1600ms
    ticker.Stop()							// ticker 打点停止
    done <- true							// 标识结束
    fmt.Println("Ticker stopped")
}