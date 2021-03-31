package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
    "unsafe"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func  String(b []byte) string {         // 字节切片转字符串（自写方法）
    return *(*string)(unsafe.Pointer(&b))
}

func main() {

    d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile("/tmp/dat1", d1, 0644)  // 写至文件 dat1
    check(err)

    f, err := os.Create("/tmp/dat2")        // 创建 dat1
    check(err)

    defer f.Close()

    d2 := []byte{115, 111, 109, 101, 10}
    fmt.Println(String(d2))
    n2, err := f.Write(d2)                  // 字节切片方式写入d2
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    n3, err := f.WriteString("writes\n")    // 字符串写入
    fmt.Printf("wrote %d bytes\n", n3)

    f.Sync()    // 调用 Sync 将缓冲区的数据写入硬盘

    w := bufio.NewWriter(f)                 // bufio 提供了的带缓冲的 Writer
    n4, err := w.WriteString("buffered\n")  // 缓冲写入字符串
    fmt.Printf("wrote %d bytes\n", n4)

    w.Flush()   // 使用 Flush 来确保，已将所有的缓冲操作应用于底层 writer

}