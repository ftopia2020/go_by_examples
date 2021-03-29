package main

import (
    "fmt"
    "math"
	"reflect"
)

const s string = "constant"

func main() {
    fmt.Println(s)

    const n = 500000000				// const 常量关键字
	fmt.Println(reflect.TypeOf(n))	// 打印类型 =》int

    const d = 3e20 / n				// e - 科学计数法
    fmt.Println(d)
    fmt.Println(reflect.TypeOf(d))	// 打印类型 =》float64

    fmt.Println(int64(d))			// 数值类型转换

    fmt.Println(math.Sin(n))        // 这里的 n 为 float64 进行计算
}