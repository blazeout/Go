package main

import (
	"fmt"
	"strings"
)

func main(){
	// \是有特殊含义的,我们应该告诉编译器我写的\就是一个单纯的\ 所以加一个\表示转义字符
	path := "\"D:\\GOfile\\src\\github.com\\day02\\07string\""
	fmt.Println(path)
	//多行的字符串
	s2 := `
	世情薄
	人情恶
	雨松黄昏花易落
`
	fmt.Println(s2)
	// len 求string类型的变量的长度
	var a  = len(s2)
	fmt.Println(a)
	// split 按照给定的字符分割
	ret := strings.Split(path,"\\")
	fmt.Println(ret)
	// contains是否包含
	// 前缀和后缀
	fmt.Println(strings.HasPrefix(path,"D"))
	fmt.Println(strings.HasSuffix(path,"\""))
	// Index判断子串出现的下标 lastindex 判断子串最后出现的位置
	s4 := "abcdefa"
	fmt.Println(strings.Index(s4,"c"))
	fmt.Println(strings.LastIndex(s4,"a"))
	fmt.Println(strings.Join(ret,"+"))
}
