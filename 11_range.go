package main

import "fmt"

func main() {

    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {		// range 迭代数组/切片，sum 累加
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {		// range 迭代数组/切片，i-下标 num-元素值
        if num == 3 {
            fmt.Println("range for 3:", i)
        }
    }

    for i, num := range nums {
        if num == 3 {
			break					// 可用 break 跳出迭代
        }
		fmt.Println("not find:", i)
    }

    for i, num := range nums {
        if num == 3 {
			continue				// 可用 continue 跳转至下一次迭代
        }
		fmt.Println("not skip:", i)
    }
	fmt.Println("--- ---")

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {			// range 迭代 map，k-key v-value
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {			// 省略 v
        fmt.Println("key:", k)
    }

    for i, c := range "go" {		// range 迭代字符串，i-字符所在位置 c-字符Unicode
        fmt.Printf("index: %d, unicode: %d, string: %q\n", i, c, c)
    }
}