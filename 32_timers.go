package main

import (
    "fmt"
    "time"
)

func main() {

	bT := time.Now()            // 开始时间
    timer1 := time.NewTimer(2 * time.Second)	// Timer定时器

    <-timer1.C		// 程序将一直阻塞，直至收到 timer1.C （相当于停顿 2s）
    fmt.Println("Timer 1 fired")

	sT := time.Since(bT)			// 从开始到当前所消耗的时间
	fmt.Println("Till now: ", sT)	// 程序运行至此约为2s
    timer2 := time.NewTimer(time.Second)	// Timer2定时器 1s
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")		// 不足2s就停啦~
    }()
    stop2 := timer2.Stop()					// Timer2停止
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

	timer3 := time.NewTimer(3 * time.Second)
	go func() {
		<-timer3.C					// 未执行完成 main()函数已经退出了
        fmt.Println("Timer 3 fired")
	}()

    time.Sleep(1 * time.Second)
	eT := time.Since(bT)			// 从开始到当前所消耗的时间
	fmt.Println("Run time: ", eT)	// 程序总共仅运行了3秒左右，协程是并发执行的
}