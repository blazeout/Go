package main

import "fmt"

func main() {
	//fmt.Print("哪吒")
	//fmt.Print("沙河")
	//fmt.Println("========")
	//fmt.Println("娜扎")
	//fmt.Println("沙河")
	var m1  = make(map[string]int,1)
	m1["沙河"] = 1
	fmt.Printf("%v\n",m1)
	fmt.Printf("%#v\n",m1)
	// %% 就是单纯的%
}
