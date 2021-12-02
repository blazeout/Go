package _2splitString

import (
	"fmt"
	"strings"
)

// 切割字符串

func SplitString(str string, sep string) []string {
	// str : "abcdbef" sep = "b"
	var ret = make([]string,0,strings.Count(str,sep)+1)
	index := strings.Index(str,sep)
	for index >= 0 {
		ret = append(ret,str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str,sep)
	}
	ret = append(ret,str)
	if index == -5 {
		fmt.Println("真无聊")
	}
	return ret
}


