package main

import (
	"os"
	"fmt"
	"flag"
	"strings"
)

func main() {

	/**
	  1. go build 要编译的文件,生成可执行文件

	  2. ./可执行文件 参数1 参数2

	 */

	argsWithProg := os.Args
	argsWithOutProg := os.Args[1:]

	fmt.Println("argsWithProg:", argsWithProg)
	fmt.Println("argsWithOutProg:", argsWithOutProg)

	arg := os.Args[3]
	fmt.Println("arg:", arg)

	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var sVar string
	flag.StringVar(&sVar, "sVar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("sVar:", sVar)

	fmt.Println("tail:", flag.Args())

	os.Setenv("nivelle", "jessy")
	fmt.Println("nivelle value:", os.Getenv("nivelle"))
	fmt.Println("not exits value:", os.Getenv("jessy"))

	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println("os.Environ()获取环境变量集合:", pair[0],"=",pair[1])
	}

	/**

	// 1. 如果省略一个标志，那么这个标志的值自动的设定为他的默认值
	// 2. flag 包需要的所有的标志出现位置参数之前（否则，这个标志将会被解析为位置参数）
nivelleMac:command nivellefu$ ./comline  -word=opt 1 2 3 4 5
argsWithProg: [./comline -word=opt 1 2 3 4 5]
argsWithOutProg: [-word=opt 1 2 3 4 5]
arg: 2
word: opt
numb: 42
fork: false
sVar: bar
tail: [1 2 3 4 5]

	*/

	/**
	nivelleMac:command nivellefu$ ./comline -h -word=temp 1 3 4 5
argsWithProg: [./comline -h -word=temp 1 3 4 5]
argsWithOutProg: [-h -word=temp 1 3 4 5]
arg: 1
Usage of ./comline:
  -fork
        a bool
  -numb int
        an int (default 42)
  -sVar string
        a string var (default "bar")
  -word string
        a string (default "foo")

	 */

}
