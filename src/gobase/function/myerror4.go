package main

import (
	"fmt"
	"os"
	"time"
)

//自定义Error
type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path,
		p.op, p.createTime, p.message)
}

func Open(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		//自定义异常对象
		return &PathError{
			path:       filename,
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}

	defer file.Close()
	return nil
}

func main() {
	err := Open("/Users/nivellefu/IdeaProjects/go-guides/README.md")
	// 获取某个实例的类型
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error,", v)
	default:

	}
}
