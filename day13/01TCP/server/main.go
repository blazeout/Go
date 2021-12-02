package main

import (
	"bufio"
	"fmt"
	"net"
)
 // 写成可以聊天的服务端和客户端
func process(conn net.Conn){
	defer conn.Close()
	for true {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil{
			fmt.Println("read from client failed , err :",err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到从client端发来的数据",recvStr)
		conn.Write([]byte(recvStr))

	}
}

func main(){
	// 1. 启动本地端口服务
	listen , err := net.Listen("tcp","127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed err :",err)
		return
	}
	// 2. 等待别人来跟我建立连接
	for true {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("start tcp server on 127.0.0.1:20000 failed err :",err)
			return
		}
		go process(conn)
	}

	//for true{
	//	conn, err := listen.Accept()
	//	if err != nil {
	//		fmt.Println("accept failed ,err:",err)
	//		continue
	//	}
	//	go process(conn)
	//}
	// 3. 与客户端通信

}