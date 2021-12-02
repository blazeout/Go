package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 打开文件写内容
// 利用file变量的write方法写
func writeDemo1(){
	fileObj, err := os.OpenFile("./xx.txt",os.O_WRONLY | os.O_CREATE | os.O_TRUNC,0644)
	if err != nil {
		fmt.Printf("open failed err : %v",err)
		return
	}
	// write
	fileObj.Write([]byte("小鸟了\n"))
	fileObj.WriteString("我他妈人傻了!")
	fileObj.Close()
}

// 利用bufio包来写文件
func writeDemo2(){
	fileObj, err := os.OpenFile("./zz.txt",os.O_WRONLY | os.O_CREATE | os.O_TRUNC,0111)
	if err != nil {
		fmt.Println("open file err")
		return
	}
	// 关闭文件
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {
			fmt.Println("关闭失败")
			return
		}
	}(fileObj)
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello world!") // 写到缓存中
	wr.Flush() // 将缓存里面的东西写到文件中
}

// 使用ioutil包来写文件

func writeDemo3(){
	err := ioutil.WriteFile("./zz.txt",[]byte("Hello Gophers"),0111)
	if err != nil {
		fmt.Printf("写入失败! err : %v\n", err)
		return
	}
}
func main() {
	writeDemo3()
}
