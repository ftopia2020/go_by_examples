package main

import "fmt"

func main() {

    var a [5]int			// 基础定义数组的方式，[长度]类型
    fmt.Println("emp:", a)	// 数组元素 =》初始值

    a[4] = 100				// 数组内元素赋值
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))	// len()方法获取数组长度

    b := [5]int{1, 2, 3, 4, 5}	// 初始化数组
    fmt.Println("dcl:", b)

    var twoD [2][3]int			// 二维数组
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}