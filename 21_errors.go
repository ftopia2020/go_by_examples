package main

import (
    "errors"
    "fmt"
)
func f1(arg int) (int, error) {		// error类型是一个内置接口类型
    if arg == 42 {
        return -1, errors.New("can't work with 42") // errors.New 构造error返回值
    }
    return arg + 3, nil             // 返回错误值为 nil, 代表没有错误
}

type argError struct {              // 构建结构体 argError
    arg  int
    prob string
}
func (e *argError) Error() string { // 对 *argError 自定义错误字符串返回
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func f2(arg int) (int, error) {     // 自定义错误类型的数据
    if arg == 42 {
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {
    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)    // 42 时，抛出错误
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}