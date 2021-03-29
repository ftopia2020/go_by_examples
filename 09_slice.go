package main

import "fmt"

func main() {

    s := make([]string, 3)      // 初始化切片，可使用 make 关键字，可初始化切片长度，填充类型初始值
    fmt.Println("emp:", s)
    i := make([]int, 3)
    fmt.Println("emp:", i)

    s[0] = "a"                  // 切片赋值同数组
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])   // 获取切片对应元素的值

    fmt.Println("len:", len(s)) // len() 内置方法获取切片长度
    fmt.Println("--- ---")

    s = append(s, "d")          // append 方法增加切片元素 append(T[], elem)
    s = append(s, "e", "f")     // 可逗号增加多个元素 append(T[], elem, elem) 
    fmt.Println(" apd:", s)
    c := make([]string, len(s))
    copy(c, s)                  // copy(target, original) 实现浅拷贝切片
    a := s                      // = 赋值为指针复制，改变原切片或新切片都会对另一个产生影响
    s[0] = "x"
    fmt.Println("copy:", c)
    fmt.Println("   a:", a)
    a[1] = "y"
    fmt.Println("   a:", a)
    fmt.Println("   s:", s)
    fmt.Println("--- ---")

    l := s[2:5]                 // 切片截断
    fmt.Println("sl1:", l)

    l = s[:5]                   // 表示从 0 -> 5
    fmt.Println("sl2:", l)

    l = s[2:]                   // 表示从 2 -> 最后
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)
    fmt.Println("--- ---")

    twoD := make([][]int, 3)    // 二维切片
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println(" 2d: ", twoD)

    cTwoD := make([][]int, 3)
    copy(cTwoD, twoD)
    fmt.Println("c2d: ", cTwoD)
    cTwoD[0][0] = 1
    fmt.Println(" 2d: ", twoD)      // 浅拷贝，twoD 的值也跟随变化
}
