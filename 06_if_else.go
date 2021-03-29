package main

import "fmt"

func main() {

    if 7%2 == 0 {					// if/else 基础条件语句
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 8%4 == 0 {					// else 可省略
        fmt.Println("8 is divisible by 4")
    }

    if num := 9; num < 0 {			// 初始;条件
        fmt.Println(num, "is negative")
    } else if num < 10 {			// else if 增加条件
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}