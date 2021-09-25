package main

import "fmt"

// 类型断言

func assign(a interface{}){
	fmt.Printf("%T\n",a)
	str,ok := a.(string)
	if !ok {
		fmt.Printf("错了!")
	}else{
		fmt.Println("传进来的是一个字符串: ",str)
	}
}

// 断言2

func assign2(a interface{}){
	fmt.Printf("%T\n" , a)
	switch t := a.(type) {
	case string:
		fmt.Println("传进来的是一个字符串",t)
	case int64:
		fmt.Println("是一个int64",t)
	case bool:
		fmt.Println("是一个布尔值",t)
	}
}

func main() {
	assign("halo")
	assign2(int64(22))
}
