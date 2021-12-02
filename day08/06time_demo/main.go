package main

import (
	"fmt"
	"time"
)

// 时间time包

// 把一个字符串的时间传换成时间粗


func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	ret := time.Unix(1624355631,0)
	fmt.Println(ret)
	fmt.Println(time.Second)
	// now 一个小时之后 是多少 now + 1
	fmt.Println(now.Add(time.Hour * 24))
	// 定时器 演示
	//timer := time.Tick(time.Second)
	//for t := range timer {
	//	fmt.Printf("%v \n",t)
	//}
	// 格式化时间 把语言中的时间对象 转换成字符串类型的时间
	// 2019.08.04
	fmt.Println(now.Format("2006-01strconv_demo-02 15-04-05"))
	fmt.Println(now.Format("2006-01strconv_demo-02 03-04-05 PM"))
	fmt.Println(now.Format("2006/01strconv_demo/02 15:04:05.000"))
	// 按照对应的格式解析字符串类型的时间
	timeObj,err := time.Parse("2006-01strconv_demo-02","2021-0 6-24")
	if err != nil {
		fmt.Println("您传入的时间格式是不对的!")
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
}
