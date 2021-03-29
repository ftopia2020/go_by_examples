package main

import "fmt"

func plus(a int, b int) int {		// func 函数方法接收 传参 & 出参
    return a + b					// 必要 return
}

func plusPlus(a, b, c int) int {
    return a + b + c
}

func printSum(a int, b int)  {				// 可无出参
	fmt.Printf("%d+%d = %d\n", a, b, a+b)	// 不用 return
}

func main() {

    res := plus(1, 2)						// 通过函数名称调用
    fmt.Println("1+2 =", res)
	printSum(1,2)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}