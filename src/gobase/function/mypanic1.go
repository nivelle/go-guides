package main

import (
	"fmt"
)

func ShowPanic1() {
	defer func() {
		fmt.Println("我虽然在异常之前，我被执行了。")
	}()
	myReCover()
	fmt.Println("上面发生了异常,被defer 捕获,我依然会执行了")
	defer func() {
		fmt.Println("我虽然在异常之后，我被执行了。")
	}()
}

func ShowPanic2() {
	/**
	  1. 利用recover处理panic指令,defer 必须放在 panic 之前定义,另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散

	  2. recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点

	  3. 多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用
	 */
	defer func() {
		err := recover() //只能捕获最近的异常
		fmt.Println("ShowPanic2 内部尝试捕获异常")
		if err != nil {
			fmt.Println("捕获了异常，异常内容2:", err)
		} else {
			fmt.Println("无内鬼2")
		}
	}()
	defer func() {
		fmt.Println("我在panic异常之前，我还能执行。")
	}()
	panic("我异常了")
	defer func() {
		fmt.Println("我在panic异常之后，我被执行不了。")
	}()
}

/**
   1. 在defer函数中，通过 recover 来终止一个goroutine的panicking过程，从而恢复正常代码的执行

   2. 可以获取通过panic传递的error
 */
func myReCover() {
	var arr [3]int
	//延迟调用匿名函数(defer 函数在主函数结束之前最后调用，可以捕获主函数中的异常)
	defer func() {
		err := recover()
		fmt.Println("myReCover 内部尝试捕获异常：")
		if err != nil {
			fmt.Println("异常内容:", err)
		} else {
			fmt.Println("无内鬼")
		}
	}()
	index := 7
	arr[index] = 9
	fmt.Println("没有异常,执行完成")
}

func main() {
	ShowPanic1()
	ShowPanic2()

}
