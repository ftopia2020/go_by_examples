package main

import "fmt"

func main() {

    i := 1
	// 单一条件，类似 while
    for i <= 3 {				// for 关键字是 Go 中唯一的循环结构
        fmt.Println(i)
        i = i + 1
    }

    for j := 7; j <= 9; j++ {	// 经典的初始;条件;后续
        fmt.Println(j)
    }

    for {						// 死循环，通常用 break、return 跳出
        fmt.Println("loop")
        break					// break 关键字跳出循环
    }

    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue			// continue 关键字直接跳转下一次迭代
        }
        fmt.Println(n)
    }
}