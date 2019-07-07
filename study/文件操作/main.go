package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func testRead() {
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println("open file faild, err：", err)
		return
	}
	defer file.Close()
	for {
		s := make([]byte, 128, 128)
		_, err := file.Read(s)
		if err == io.EOF {
			fmt.Println("文件读完了")
			return
		}
		if err != nil {
			fmt.Println("read from filr faild, err：", err)
			return
		}
		fmt.Println(string(s))
	}

}

func testBufIo() {
	file, err := os.Open("./hello.txt")
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	reader := bufio.NewReader(file)
	defer file.Close()

	// 也存在问题
	for {
		// \n截止字符
		content, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print("文件读取完毕\n")
			return
		}
		if err != nil {
			fmt.Print("打开文件失败")
			return
		}
		fmt.Print(string(content))
	}

}

func testIoUtil() {
	content, err := ioutil.ReadFile("./hello.txt")
	if err != nil {
		fmt.Println("文件打开失败，err：", err)
	}
	fmt.Println(string(content))

}

func testWrite1() {
	file, err := os.OpenFile("./write1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("写入失败，err：", err)
		return
	}
	defer file.Close()
	fmt.Println(file)
	// 字符串方式写入
	_, _ = file.WriteString("hello python\n")
	// 字节切片方式写入
	_, _ = file.Write([]byte("hello go\n"))

}

func testWrite2() {
	err := ioutil.WriteFile("write2.txt", []byte("hello world\n"), 0777)
	if err != nil {
		fmt.Println("写入出错，err：", err)
	}

}

func main() {

	// os.Open() 不可控，有异常
	//testRead()

	// bufio
	//testBufIo()

	// ioutil：直接传入文件路径
	//testIoUtil()

	// write1
	//testWrite1()

	// write2：模式默认为清空
	testWrite2()
}
