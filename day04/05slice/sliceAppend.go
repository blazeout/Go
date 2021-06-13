package main

import "fmt"

func main(){
	s1 := []string{"北京","上海","深圳"}
	s1 = append(s1, "广州")
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(s1)
	s1 = append(s1,"杭州","成都")
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	ss := []string{"长沙","武汉"}
	// ss... 表示拆开元素
	// cap 乘以 2 若是 容量小于1024的话
	s1 = append(s1,ss...)
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}
