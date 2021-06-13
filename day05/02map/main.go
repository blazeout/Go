package main

import "fmt"

func main() {
	// 要估计map的容量,避免动态扩容
	var m1  = make(map[string]int,10)
	m1["理想"] = 18
	m1["小王"] = 19
	fmt.Println(m1)
	fmt.Println(m1["理想"])
	value,ok := m1["李选哪个"]
	if !ok{
		fmt.Printf("查无次key\n")
	}else{
		fmt.Println(value)
	}

	// map的遍历
	for k,v := range m1{
		fmt.Println(k,v)
	}
	// 删除
	delete(m1,"理想")
	fmt.Println(m1)
}
