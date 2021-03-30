package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {	// worker 函数
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {

	bT := time.Now()            // 开始时间
    const numJobs = 5			// 缓冲 5 个 worker
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    for w := 1; w <= 3; w++ {	// 实际遍历 1, 2, 3
        go worker(w, jobs, results)
    }

    for j := 1; j <= numJobs; j++ {
        jobs <- j				// 遍历写入值至 jobs
    }
    close(jobs)					// 关闭 jobs 通道

    for a := 1; a <= numJobs; a++ {
         res := <-results		// 接收 results 值
		 fmt.Println(res)		// j*2
    }
	eT := time.Since(bT)			// 从开始到当前所消耗的时间
	fmt.Println("Run time: ", eT)	// 程序总共仅运行了约2秒，协程是并发执行的
}