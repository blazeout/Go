package main

import "fmt"

func main(){
	// 字符串修改
	str := "白萝卜"// string类型
	s2 := []rune(str)
	// 把字符串强制转化为一个rune切片
	s2[0] = '红'// rune类型实际为int32类型
	fmt.Println(string(s2))


	// 类型转换
	n := 10
	var f float64
	f = (float64(n))
	fmt.Println(f)
	fmt.Printf("%T",f)
}
