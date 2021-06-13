package main

import "fmt"

// 运算符
func main(){
	var(
		a = 5
		b = 2
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a * b)
	// ++和-- 在go语言中是一个语句不能放在等号的右边赋值
	a++
	b--
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	// go语言是强类型的,相同类型才能比较
	var year int = 20
	if year > 18 && year < 60{
		fmt.Printf("年龄是%d",year)
	}else{
		fmt.Println("不用上班")
	}
	if year < 18 || year > 60{
		fmt.Println("不要上班")
	}else{
		fmt.Println("苦逼上班")
	}
	// 逻辑取反
	isInt := false
	fmt.Println(isInt)
	fmt.Println(!isInt)

	// 位运算符 针对的是二进制
	// 5的二进制是0101
	// 2的二进制是0010
	// &:按位与运算 有0则为0
	// 预计结果为 0
	fmt.Println(5 & 2)
	// | : 按位或运算,有1则为1
	// 预计结果为 7
	fmt.Println(5 | 2)
	// ^ 异或 : 按位异或 两位不一样则为1
	// 预计为 7
	fmt.Println(5 ^ 2)
	// >> : 右移运算符 把所有的二进制为向右移x个单位 格式为 变量 >> x
	// 预计为1
	fmt.Println(6 >> 2)
	// << : 左移运算符 把所有的二进制位向左移x个单位
	fmt.Println(6 << 2)


}
