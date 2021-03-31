package main

import (
    "bufio"
    "fmt"
    "net/http"
)

func main() {

	// 向服务端发送一个 HTTP GET 请求
    resp, err := http.Get("http://gobyexample.com")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

	// 打印 HTTP response 状态 200 OK
    fmt.Println("Response status:", resp.Status)

    scanner := bufio.NewScanner(resp.Body)
    for i := 0; scanner.Scan() && i < 5; i++ {	// 打印 response body 前面 5 行的内容
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}