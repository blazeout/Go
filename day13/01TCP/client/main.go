package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client 端

func main() {

	// 1. 与server端建立连接
	conn, err := net.Dial("tcp","127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial address failed err:",err)
		return
	}
	// 2. 通信
	reader := bufio.NewReader(os.Stdin)
	for true{
		fmt.Print("请输入要输入的内容")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit"{
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}
