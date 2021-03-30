package main

import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)

func main() {

    var state = make(map[int]int)		// 记录 state 于 map 中

    var mutex = &sync.Mutex{}			// 定义互斥锁

    var readOps uint64
    var writeOps uint64

    for r := 0; r < 100; r++ {			// 100个协程模拟读操作
        go func() {
            total := 0
            for {
                key := rand.Intn(5)		// 生成 1~5 随机数
                mutex.Lock()			// +锁
                total += state[key]		// 重复读取 state
                mutex.Unlock()			// 解锁
                atomic.AddUint64(&readOps, 1)	// 原子计数
                time.Sleep(time.Millisecond)	// ms级停顿
            }
        }()
    }
    time.Sleep(time.Second)
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
	fmt.Println("--- ---")

    for w := 0; w < 10; w++ {			// 10个协程模拟写操作
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)

    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}