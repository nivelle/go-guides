package main

import (
	"errors"
	"fmt"
)
//导致关键流程出现不可修复性错误的使用 panic，其他使用 error
func Dev(dividend int, divisor int) (quotient int, err error) {
	if divisor == 0 {
		err = errors.New("除数不能为0")
		return
	} else {
		quotient = dividend / divisor //将结果赋值给定义好的返回值
		return                        // err默认nil
	}

}

func main() {
	result, err := Dev(2, 0)
	fmt.Println(result)
	fmt.Println(err)
}
