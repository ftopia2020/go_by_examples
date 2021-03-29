package main

import "fmt"

func fact(n int) int {
    if n == 0 {
        return 1
    }
	if n > 1 {
		fmt.Printf("%d*", n)
	} else {
		fmt.Printf("1")
	}
    return n * fact(n-1)	// 支持递归，即自己调用自己
}

func main() {
    fmt.Printf("=%d\n", fact(7))
}