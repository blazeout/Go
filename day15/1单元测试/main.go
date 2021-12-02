package main

import (
	"fmt"
	"github.com/day15/02splitString"
)

func main() {
	str := _2splitString.SplitString("babcdbef","b")
	fmt.Printf("%#v\n",str)
	str1 := _2splitString.SplitString("bbb","b")
	fmt.Printf("%#v\n",str1)
	str2 := _2splitString.SplitString("bcb","b")
	fmt.Printf("%#v\n",str2)
}