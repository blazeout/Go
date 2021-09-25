package calc

// Add 包中的标识符如果首字母是小写的,表示是私有的,只能在这个包中使用
// Add 首字母大写的标识符,可以被外部的包调用
func Add(x ,y int)int{
	return x + y
}
