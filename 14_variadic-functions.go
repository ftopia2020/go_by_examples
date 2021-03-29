package main

import "fmt"

func sum(nums ...int) {			// ...dataType 不定长参数，表示接收任意数量该类型的传参
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}	
    sum(nums...)				// 切片...，类似语法糖
}