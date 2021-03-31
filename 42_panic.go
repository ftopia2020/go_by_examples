package main

import (
	"os"
	"fmt"
) 

func main() {

    // panic("a problem")		// 直接抛出异常，后续代码不会执行

    _, err := os.Create("/tmp/file")
    if err != nil {			// 收集错误信息，抛出异常
        panic(err)
    }

	fmt.Println("OK")
}