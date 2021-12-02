package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("listen failed err,", err)
		return
	}
	defer func(listen *net.UDPConn) {
		err := listen.Close()
		if err != nil {
		}
	}(listen)
	var data [1024]byte
	for true {
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed err:", err)
			continue
		}
		str := strings.ToUpper(string(data[:n]))
		_, err = listen.WriteToUDP([]byte(str), addr)
		if err != nil {
			continue
		}
	}
}
