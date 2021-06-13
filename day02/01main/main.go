package _1main
import "fmt"
// 单独声明
//var isOK bool
// 批量声明

var(
	name string
	age int
	isOk bool
	studentHeight int

)
func main(){
	name = "wangjiahao"
	age = 18
	var heHe string
	heHe = "name"
	fmt.Println(heHe)
	isOk = true
	// GO语言中变量声明必须使用,不使用就编译不过去
	fmt.Printf("name = %s",name)
	fmt.Println(age)// 打印完内容 会换行
	fmt.Print(isOk)
	fmt.Println(heHe)
	var s1  = "laowang"
	var s2 = "is20sui"
	fmt.Println(s1,s2)
	// 简短变量声明
	s3 := "hhhh"
	fmt.Println(s3)
	// 同一个作用域中不能重复声明一个变量

}

