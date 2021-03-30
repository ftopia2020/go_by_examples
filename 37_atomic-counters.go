package main

import (
    "fmt"
    "sync"
    "sync/atomic"
    "time"
)

func main() {

    bT := time.Now()            // 开始时间
    var ops uint64

    var wg sync.WaitGroup

    for i := 0; i < 50; i++ {   // 50个协程
        wg.Add(1)

        go func() {
            for c := 0; c < 1000; c++ {

                atomic.AddUint64(&ops, 1)
            }
            wg.Done()
        }()
    }

    wg.Wait()          // 等待所有协程

    fmt.Println("ops:", ops)

    eT := time.Since(bT)			// 从开始到当前所消耗的时间
	fmt.Println("Run time: ", eT)	// 程序总共仅运行了约1秒，协程是并发执行的
}