package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter,r *http.Request){
	str, err := ioutil.ReadFile("D:\\GOfile\\src\\github.com\\day14\\01http\\server\\tt.txt")
	if err != nil {
		w.Write([]byte("暂无次内容"))
		return
	}
	w.Write(str)
}

func f2(w http.ResponseWriter, r *http.Request){
	// 对于GET请求,参数都放在了URL上面,请求体中是没有数据的
	queryParam := r.URL.Query()
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name,age)

	fmt.Println(r.Proto)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/posts/Go/15_socket/",f1)
	http.HandleFunc("/xxx/",f2)
	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		return
	}

}
