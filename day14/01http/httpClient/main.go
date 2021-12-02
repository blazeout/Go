package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp , err := http.Get("http://127.0.0.1:9090/xxx/?name=sb&age=18")
	if err != nil{
		fmt.Printf("get url failed err: %v",err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read resp.Body failed err :",err)
		return
	}
	fmt.Println(string(b))
}
