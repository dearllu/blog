package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm() //解析参数
	fmt.Println(r.Form) //信息输出在服务器段
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form {
		fmt.Println("key",k)
		fmt.Println("val: ",strings.Join(v,""))
	}
	fmt.Println(w,"hello astaxie!")
}

func main() {
	http.HandleFunc("/",sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal("listen and serve: ",err)
	}
}