package main

import (
    "fmt"
    "sort"
)

type byLength []string			// 定义数组类型-按长度排序的数组

// 为该类型实现了 sort.Interface 接口的 Len、Less 和 Swap 方法（自定义）
func (s byLength) Len() int {			
    return len(s)
}

func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {	// 判断排序为字符串长度短的在前
    return len(s[i]) < len(s[j])
}

func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}