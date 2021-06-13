package main

import "fmt"

func main() {
	// map 和 slice 的结合
	// 元素类型为map的切片
	var s1 = make([]map[int]string,5,10)
	s1[0] = make(map[int]string,1)
	// 指取到的s1的第0个元素map
	s1[0][10] = "a"
	fmt.Println(s1)
	// 值为切片类型的map
	var m1 = make(map[string][]int,5)
	m1["北京"] = []int{10,20,30}
	m1["上海"] = make([]int,5,5)
	m1["上海"] = []int{10,2030}
	fmt.Println(m1)
}
