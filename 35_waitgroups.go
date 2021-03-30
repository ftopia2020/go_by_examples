package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {	// 注意，WaitGroup 必须通过指针传递给函数。

	// defer 可理解为后入先出的栈，这边先执行后续代码，再执行 wg.Done
    defer wg.Done()				// return时，通知 WaitGroup，当前协程的工作已经完成

    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)		// 停顿 1s，模拟耗时的任务
    fmt.Printf("Worker %d done\n", id)
}

func main() {

	bT := time.Now()            // 开始时间
    var wg sync.WaitGroup		// 等待多个协程完成，我们可以使用 WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

	time.Sleep(500 * time.Millisecond)
    wg.Wait()						// 阻塞，直到 WaitGroup 计数器恢复为 0； 即所有协程的工作都已经完成
	eT := time.Since(bT)			// 从开始到当前所消耗的时间
	fmt.Println("Run time: ", eT)	// 程序总共仅运行了约1秒，协程是并发执行的
}