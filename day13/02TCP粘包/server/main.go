package main

import (
	"bufio"
	"fmt"
	protocol "github.com/day13/02TCP粘包/nianbaosolution"
	"io"
	"net"
)

func main(){
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed err",err)
		return
	}
	defer listen.Close()
	for true {
		conn , err := listen.Accept()
		if err != nil{
			fmt.Println("accept failed err:",err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for{
		receiveStr, err := protocol.Decode(reader)
		if err == io.EOF {
			fmt.Println("data have been read none")
			break
		}
		if err != nil{
			fmt.Println("read from client failed err:",err)
			break
		}
		fmt.Println("收到从客户端发来的消息 :",receiveStr)
	}
}