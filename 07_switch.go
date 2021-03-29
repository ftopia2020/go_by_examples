package main

import (
    "fmt"
    "time"
)

func main() {

    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {		// 基础的多条件 switch 结构
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:	// 逗号,隔开 支持多值 =》同一分支逻辑
        fmt.Println("It's the weekend")
    default:							// 可选 default 分支
        fmt.Println("It's a weekday")
    }

    t := time.Now()
    switch {							// 不带表达式的 switch，是实现 if/else 逻辑的另一种方式
    case t.Hour() < 12:					// case 为条件表达式
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    whatAmI := func(i interface{}) {	// switch + func + interface 判断变量类型的写法
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}