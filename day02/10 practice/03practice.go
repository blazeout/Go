package main

import (
	"fmt"
	"unicode"
)

// 编写代码统计出字符串"hello沙河小王子"中的汉字数量
func main(){

 str2 := "hello沙河小王子"
 count := 0
	for _, i2 := range str2 {
		if unicode.Is(unicode.Han ,i2){
			count++
		}
	}
	fmt.Println(count)
}
