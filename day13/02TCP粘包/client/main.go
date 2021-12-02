package main

import (
	"fmt"
	protocol "github.com/day13/02TCP粘包/nianbaosolution"
	"net"
)

func main(){
	conn, err := net.Dial("tcp","127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed err:",err)
		return
	}
	defer conn.Close()
	for i := 0; i < 10; i++ {
		msg := "hello how are you?"
		encodeMsg, err := protocol.Encode(msg)
		if err != nil {
			fmt.Println("encode failed err : ",err)
			return
		}
		conn.Write(encodeMsg)
	}
}
