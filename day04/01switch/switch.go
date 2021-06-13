package main

import "fmt"

func main(){
	//var finger int
	var(
		n int


	)
	fmt.Println("请输入数字")
	//fmt.Scanf("%d",&finger)
	//switch finger {
	//case 1:
	//	fmt.Println("大拇指")
	//case 2:
	//	fmt.Println("食指")
	//case 3:
	//	fmt.Println("中指")
	//case 4:
	//	fmt.Println("无名指")
	//case 5:
	//	fmt.Println("小拇指")
	//default :
	//	fmt.Println("错误的输入")
	fmt.Scanf("%d",&n)
	switch n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println("啥也不是")
	}
		age := 30
		switch {
		case age < 25:
			fmt.Println("好好学习吧")
		case age > 25 && age < 35:
			fmt.Println("好好工作吧")
			fallthrough
		case age > 60:
			fmt.Println("好好享受吧")
		default:
			fmt.Println("活着真好")
		}
	}

