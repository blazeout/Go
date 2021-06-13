package main

import "fmt"

// 编写一件代码分别定义整形,浮点型,布尔型,字符串型变量
// 使用fmt.println()搭配%T 分别打印出上述变量的值和类型
var(
	a int
	b float64
	c bool
	str1 string
)
func main(){
	a = 1
	b = 0.0
	c = true
	str1 = "fuckboy"
fmt.Printf("%d %T\n",a,a)
fmt.Printf("%f %T\n",b,b)
fmt.Printf("%T\n",c)
fmt.Printf("%s %T\n",str1,str1)
}
