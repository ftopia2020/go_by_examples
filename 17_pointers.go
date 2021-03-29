package main

import "fmt"

func zeroval(ival int) {
    ival = 0					// 相当于变量赋值为 0
}

func zeroptr(iptr *int) {		// *int，使用了int指针
    *iptr = 0					// 通过指针赋值0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)					// 通过 &i 语法来取得 i 的内存地址，即指向 i 的指针
    fmt.Println("zeroptr:", i)  // 通过指针赋值给 i

    fmt.Println("pointer:", &i)	// 指针地址可打印
}