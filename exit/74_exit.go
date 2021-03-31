package main

import (
    "fmt"
    "os"
)

func main() {

	// 当使用 os.Exit 时 defer 将不会 被执行， 所以这里的 fmt.Println 将永远不会被调用
    defer fmt.Println("!")

    os.Exit(3)  // 退出并且退出状态为 3
}