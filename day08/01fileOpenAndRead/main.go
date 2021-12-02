package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 打开文件

func readFromFile(){
	fileObj, err := os.Open("./xx.txt")
	if err != nil{
		fmt.Printf("open file failed error : %v\n", err)
		return
	}
	// 记得关闭文件
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {
			return
		}
	}(fileObj)
	var tmp = make([]byte, 1024)
	n, err := fileObj.Read(tmp)
	if err != nil{
		fmt.Printf("read from file failed , err :%v",err)
		return
	}
	fmt.Println("读了",n,"个字节")
	fmt.Println(string(tmp[:n]))
}

// 利用bufio这个包读文件
func readFromFileByBuff(){
	fileObj, err := os.Open("./xx.txt")
	if err != nil{
		fmt.Printf("open file failed error : %v\n", err)
		return
	}
	// 记得关闭文件
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {
			return
		}
	}(fileObj)
	reader := bufio.NewReader(fileObj)
	for true {
		line,err := reader.ReadString('\n')
		if err == io.EOF{
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Printf("read from file failed err : %v",err)
			return
		}
		fmt.Print(line)
	}
}

func readFromFileByIoutil(){

	ret, err :=ioutil.ReadFile("./xx.txt")
	if err != nil {
		fmt.Println("read file failed!")
		return
	}
	fmt.Println(string(ret))
}

func main(){
	//readFromFileByBuff()
	readFromFileByIoutil()
}
