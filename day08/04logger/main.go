package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileObj, err := os.OpenFile("./xx.log",os.O_WRONLY | os.O_CREATE | os.O_APPEND,0644)
	if err != nil {
		fmt.Println("open file failed")
		return
	}
	//log.SetOutput(os.Stdout)
	// 设置输出目的地
	log.SetOutput(fileObj)
	for  {
		log.Println("这是一条测试的日志")
		time.Sleep(time.Second * 3)
	}
}
