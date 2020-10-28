package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/errors"
	"fmt"
)

//标准库 errors.New 和 fmt.Errorf 函数用于创建实现 error 接口的错误对象
var ErrDivByZero = errors.New("division bu zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}

func main() {

	defer func() {
		fmt.Println(recover())
	}()
	z, err := div(10, 1)
	switch err {
	case nil:
		println("计算结果:", z)
	case ErrDivByZero:
		panic(err)
	}
}
