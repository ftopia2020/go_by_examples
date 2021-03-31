package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // 最基本的文件读取任务，将文件内容读取到内存中
    dat, err := ioutil.ReadFile("/tmp/dat1")
    check(err)
    fmt.Print(string(dat))  
    fmt.Println("--- ---")

    f, err := os.Open("/tmp/dat1")
    check(err)
    defer f.Close()         // 注意一旦打开，就须要定义 close()

    b1 := make([]byte, 5)   // 5个字节切片
    n1, err := f.Read(b1)   // 读取并记录于n1
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

    o2, err := f.Seek(6, 0) // 光标位置定位
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: ", n2, o2)
    fmt.Printf("%v\n", string(b2[:n2]))

    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2) // io 包提供了一个更健壮的实现 ReadAtLeast
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    _, err = f.Seek(0, 0)
    check(err)
    r4 := bufio.NewReader(f)    // bufio 包实现了一个缓冲读取器；提供了很多附加的读取函数
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))


}