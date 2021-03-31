package main

import (
	"fmt"
    "os"
    "os/exec"
    "syscall"
)

func main() {

	/*
		我们将执行 ls 命令。 Go 要求我们提供想要执行的可执行文件的绝对路径， 
		所以我们将使用 exec.LookPath 找到它（应该是 /bin/ls）
	*/
    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }
	fmt.Println(binary)

    args := []string{"ls", "-a", "-l", "-h"}	// Exec 需要的参数是切片的形式的

    env := os.Environ()	// Exec 同样需要使用环境变量。 这里我们仅提供当前的环境变量

	/*
		这里是真正的 syscall.Exec 调用。 如果这个调用成功，那么我们的进程将在这里结束，并被 /bin/ls -a -l -h 进程代替。 
		如果存在错误，那么我们将会得到一个返回值。
	*/
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}