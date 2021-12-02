package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main(){
	socket, err := net.DialUDP("udp",nil,&net.UDPAddr{
		IP: net.IPv4(127,0,0,1),
		Port: 40000,
	})
	if err != nil {
		return
	}
	defer socket.Close()
	reader := bufio.NewReader(os.Stdin)
	data := make([]byte,1024)
	for true {
		fmt.Println("请输入要发送的内容")
		sendData , _:= reader.ReadString('\n')
		_, err = socket.Write([]byte(sendData))
		if err != nil {
			fmt.Println("发送数据失败 err : ",err)
			return
		}
		_, _, err = socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("接受数据失败 err : ",err)
			return
		}
		fmt.Print("收到从服务端回复的数据:")
		fmt.Println(string(data))
	}
}
