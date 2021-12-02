package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj , err := os.OpenFile("./xx.txt",os.O_RDWR | os.O_CREATE,0644)
	if err != nil{
		fmt.Println("文件打开失败")
		return
	}
	defer func(fileObj *os.File) {
		err := fileObj.Close()
		if err != nil {
			return
		}
	}(fileObj)
	_, err = fileObj.Seek(1, 0)
	if err != nil {
		fmt.Printf("err : %v",err)
		return
	}
	var s []byte
	s = []byte{'c'}
	_, err = fileObj.Write(s)

	var ret [1]byte
	read, err := fileObj.Read(ret[:])
	if err != nil {
		return
	}
	fmt.Println(string(ret[:read]))
}
