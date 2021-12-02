package main

import (
	"fmt"
	"reflect"
)

type person struct{
	name string
	id int
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func reflectType(v interface{}){
	x := reflect.TypeOf(v)
	fmt.Printf("type: %v\n",x)
	fmt.Printf("type name : %v type kind : %v\n",x.Name(),x.Kind())
}

// 利用反射修改值 需要传指针
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200) // 需要使用Elem()方法来获取指针对应的值
	}
}

func main() {
	var a int64 = 3
	reflectType(a)
	var b float64 = 100.000
	reflectType(b)
	var p  = person{}
	reflectType(p)
	var f float64 = 3.14
	reflectValue(f)
	reflectSetValue1(&a)
}
