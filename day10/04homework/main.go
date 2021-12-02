package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int	`ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host string `ini:"host"`
	Port int `ini:"port"`
	Password string `ini:"password"`
	Database int `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

// 要实现的函数 利用反射将ini文件里面的信息读取到mc结构体中
func loadIni(fileName string,data interface{})(err error){
	// 0.参数的校验 传进来的data参数必须是指针类型 因为需要在函数中对其赋值
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err := errors.New("data param should be a pointer")
		return err
	}
	// 0.1 传进来的data参数必须是结构体类型(配置文件中各种键值对需要赋值给结构体的字段)
	// elem 指针指向的值
	if t.Elem().Kind() != reflect.Struct {
		err := errors.New("data param should be a struct pointer")
		return err
	}
	// 1. 读文件得到字节类型的数据
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	// 2.一行一行的读取数据
	// 将文件内容转换为字符串 将字符串按照换行符分割
	lineSlice := strings.Split(string(file),"\r\n")
	var structName string
	for index, line := range lineSlice {
		// 去掉每个S的首尾空格
		line =  strings.TrimSpace(line)

		// 2.1 如果是注释 就跳过
		if strings.HasPrefix(line,";") || strings.HasPrefix(line,"#"){
			continue
		}
		// 2.2 如果是[]开头的就表示是节(section)
		if strings.HasPrefix(line,"[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err := fmt.Errorf("line : %d syntax error", index+1)
				fmt.Println(err)
				return err
			}
				// 把这一行首尾的[]去掉 然后去掉空格最后取到中间的内容 判断是否为0
				sectionName := strings.TrimSpace(line[1:len(line)-1])
				if len(sectionName) == 0 {
					err := fmt.Errorf("index : %d [] contains valid neiron",index + 1)
					return err
				}
				// 根据sectionName 去 data里面找对应的结构体

				for  i := 0; i < t.Elem().NumField() ; i++ {
					field := t.Elem().Field(i)
					filedName := field.Tag.Get("ini")
					// 找到了对应的结构体
					if sectionName == filedName {
						structName = field.Name
						fmt.Printf("找到了%s对应的嵌套结构体为%s\n",filedName,structName)
					}
				}
		} else {

			// 2.3 如果不是[开头的就是 = 分割的键值对
			// 1. 以等号分割这一行,等号左边是key , 等号右边是value
			if strings.Index(line,"=") == -1 || strings.HasPrefix(line,"="){
				err = fmt.Errorf("line : %d syntax error",index + 1)
				return err
			}
			idx := strings.Index(line,"=")
			key := strings.TrimSpace(line[:idx])
			//value := strings.TrimSpace(line[idx+1:])
			// 2.根据stcuctName 去 data 里面把对应的嵌套结构体找出来
			sValue := reflect.ValueOf(data)
			sType := sValue.Type()
			structObj := sValue.Elem().FieldByName(structName)
			if structObj.Kind() != reflect.Struct {
				err = fmt.Errorf("data的%s字段应该是一个结构体",structName)
				return err
			}
			// 3.遍历嵌套结构体的每一个字段,判断tag是不是等于key
			var fieldName string
			for i := 0; i < structObj.NumField(); i++ {
				field := sType.Elem().Field(i)
				if field.Tag.Get("ini") == key {
					fieldName = field.Name
				}
			}
			// 4.如果是的那么就给这个字段的value赋值
			fileObj := sValue.FieldByName(fieldName)
			fmt.Println(fieldName,fileObj.Type().Kind())
		}


	}

	return
}

func main() {
	var mf Config
	//var x  = new(int)
	err := loadIni("./conf.ini",&mf)
	if err != nil{
		fmt.Printf("load ini failed err : %v\n",err)
	}
	//fmt.Println(mc.Address,mc.Port,mc.Username,mc.Password)
}
