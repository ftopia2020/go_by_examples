package main

import (
    "fmt"
    "sort"	// 内置包sort：https://golang.org/pkg/sort/
)

func main() {

    strs := []string{"c", "a", "b"}
    sort.Strings(strs)				// 字符串字母排序 大写>小写
    fmt.Println("Strings:", strs)

    ints := []int{7, 2, 4}
    sort.Ints(ints)					// 整形正序排序
    fmt.Println("Ints:   ", ints)

    s := sort.IntsAreSorted(ints)	// 判断是否排序
    fmt.Println("Sorted: ", s)
}