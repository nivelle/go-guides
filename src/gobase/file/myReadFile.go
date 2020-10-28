package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"io"
	"bufio"
)

func FileCheck(e error) {
	if e != nil {
		panic(e)
	}
}
func OpenFile() {

	/**

	  os.O_CREATE|os.O_APPEND
	  或者 os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	  os.O_RDONLY // 只读
	  os.O_WRONLY // 只写
	  os.O_RDWR // 读写
	  os.O_APPEND // 追加（Append）
	  os.O_CREATE // 如果文件不存在则先创建
	  os.O_TRUNC // 文件打开时裁剪文件
	  os.O_EXCL // 和O_CREATE一起使用，文件不能存在
	  os.O_SYNC // 以同步I/O的方式打开
	第三个参数：权限(rwx:0-7)
	  0：没有任何权限
	  1：执行权限
	  2：写权限
	  3：写权限和执行权限
	  4：读权限
	  5：读权限和执行权限
	  6：读权限和写权限
	  7：读权限，写权限，执行权限
	*/

	fp, err := os.OpenFile("/Users/nivellefu/IdeaProjects/go-guides/README.md", os.O_APPEND|os.O_WRONLY, 6) //读写方式打开
	if err != nil {
		fmt.Printf("异常内容1,%s", err)
	}
	defer func() {
		err := recover() //只能捕获最近的异常
		if err != nil {
			fmt.Printf("异常内容1,%s", err)
		}
	}()

	defer fp.Close()

	count, writErr := fp.WriteString("go->")
	if writErr != nil {
		fmt.Println("写入文件失败。", count, writErr)
		return
	}
	fmt.Println("写入字节数:", count)

	dat, err := ioutil.ReadFile("/Users/nivellefu/IdeaProjects/go-guides/README.md")
	FileCheck(err)
	fmt.Println(string(dat))
	fmt.Println("文件测试 ｜")
	f, err := os.Open("/Users/nivellefu/IdeaProjects/go-guides/README.md")
	defer f.Close()
	//最多读取2个字节
	b1 := make([]byte, 2)
	n1, err := f.Read(b1)
	fmt.Printf("最多读2个字节： %d bytes: %s \n", n1, string(b1))

	//从已知位置开始读取文件(从当前指针开始偏移2个位移，返回定位到的偏移量位置)
	o3, err := f.Seek(2, io.SeekCurrent)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	fmt.Printf("最少读%d 字节, 开始位置: %d, 读取到的内容:%s \n", n3, o3, string(b3))

	//没有内置的回转支持，但是使用 Seek(0, 0) 实现。
	_, err = f.Seek(0, io.SeekStart)
	FileCheck(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(10)
	fmt.Println("通过缓存读取10个字节:", string(b4))
}

func main() {
	OpenFile()
}
