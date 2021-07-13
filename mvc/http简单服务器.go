package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/",sayhelloName1)//设置访问的路由
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("ListenAnd Server: ",err)
	}
}



func sayhelloName1(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form {
		fmt.Println("key: ",k)
		fmt.Println("val: ",strings.Join(v,""))
	}
	fmt.Fprintf(w,"hello astaxie!!!\n")
	fmt.Fprintf(w,"name:  tom")
}