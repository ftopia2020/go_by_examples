package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {

    sigs := make(chan os.Signal, 1)		// 通过向一个通道发送 os.Signal 值来发送信号通知
    done := make(chan bool, 1)			// 程序结束时发送通知的通道

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)	// signal.Notify 注册给定的通道，用于接收特定信号

	// 协程执行一个阻塞的信号接收操作。 当它接收到一个值时，它将打印这个值，然后通知程序可以退出。
    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

	/*
		当我们运行这个程序时，它将一直等待一个信号。 通过 ctrl-C（终端显示为 ^C），
		我们可以发送一个 SIGINT 信号， 这会使程序打印 interrupt 然后退出
	*/
    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}