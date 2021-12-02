package main

import (
	"fmt"
	"strconv"
)

// strconv

func main(){
	// 从字符串里面解析出数字
	str := "10000"
	parseInt, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return
	}
	fmt.Printf("%d %T",parseInt,int32(parseInt))
	i1, err := strconv.Atoi(str)
	if err != nil {
		return
	}
	fmt.Println(i1)
	// 将一个int类型的数字转换为字符串类型
	// 1.使用itoa
	i2 := 200
	fmt.Println(i2)
	str1 := strconv.Itoa(200)
	fmt.Printf("%#v\n",str1)
	// 2.使用Sprintf返回一个字符串并且使用ret接受
	i := int32(1638)
	//ret := string(i)
	//fmt.Println(ret)
	ret  := fmt.Sprintf("%d", i)
	fmt.Printf("%#v\n",ret)
}
