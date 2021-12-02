package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main(){
	str := `"name": "周林", "age" : 9000`
	var p1 person
	err := json.Unmarshal([]byte(str), &p1)
	if err != nil {
		return
	}
	fmt.Printf("%v",p1)
}
