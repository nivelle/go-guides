package main

import (
	"os/exec"
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
)

func main() {
	dateCmd := exec.Command("date")
	//处理运行一个命令，它等待命令运行完成，并收集命令的输出。如果没有问题，dateOut 将获取到日期信息的字节
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	grepCmd.Start()
	grepIn.Write([]byte("hello grep \n goodbye grep hello"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	// -a :显示隐藏文件；-l: 列表方式显示信息 ； -h: 显示文件大小
	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()
    // syscall : 第一个参数是 可执行文件的绝对路径，第二个参数是切片形式的参数，第三个参数是环境变量
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
