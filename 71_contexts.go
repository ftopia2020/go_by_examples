package main

import (
    "fmt"
    "io/ioutil"
    "os/exec"
)

func main() {

    dateCmd := exec.Command("date")		// 执行 date 命令
    dateOut, err := dateCmd.Output()	// .Output 是另一个帮助函数，常用于处理运行命令、等待命令完成并收集其输出
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))		// 打印命令结果

    grepCmd := exec.Command("grep", "hello")
    grepIn, _ := grepCmd.StdinPipe()	// 将从外部进程的 stdin 输入数据并从 stdout 收集结果
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd.Wait()

    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

	/*
		注意，在生成命令时，我们需要提供一个明确描述命令和参数的数组，而不能只传递一个命令行字符串。 
		如果你想使用一个字符串生成一个完整的命令，那么你可以使用 bash 命令的 -c 选项
    */
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))

}