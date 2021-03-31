package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    err := os.Mkdir("subdir", 0755)
    check(err)
	// 创建这个临时目录后，一个好习惯是：使用 defer 删除这个目录。 
    defer os.RemoveAll("subdir")	// os.RemoveAll 会删除整个目录（类似于 rm -rf 慎用！）
	
    createEmptyFile := func(name string) {
        d := []byte("")
        check(ioutil.WriteFile(name, d, 0644))
    }

    createEmptyFile("subdir/file1")
	// 创建一个有层级的目录，使用 MkdirAll 函数，并包含其父目录。 这个类似于命令 mkdir -p
    err = os.MkdirAll("subdir/parent/child", 0755)
    check(err)
    createEmptyFile("subdir/parent/file2")
    createEmptyFile("subdir/parent/file3")
    createEmptyFile("subdir/parent/child/file4")

    c, err := ioutil.ReadDir("subdir/parent")
    check(err)
    fmt.Println("Listing subdir/parent")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    err = os.Chdir("subdir/parent/child")
    check(err)

    c, err = ioutil.ReadDir(".") // ReadDir 列出目录的内容，返回一个 os.FileInfo 类型的切片对象
    check(err)

    fmt.Println("Listing subdir/parent/child")
    for _, entry := range c {
        fmt.Println(" ", entry.Name(), entry.IsDir())
    }

    err = os.Chdir("../../..")  // Chdir 可以修改当前工作目录，类似于 cd
    check(err)

    fmt.Println("Visiting subdir")
    err = filepath.Walk("subdir", visit) // Walk遍历一个目录及其所有子目录，接受一个路径和回调函数（用于处理访问到的每个目录和文件）
}

func visit(p string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(" ", p, info.IsDir())
    return nil
}