package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

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

func login(w http.ResponseWriter,r *http.Request) {
	fmt.Println("method: ",r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w,nil))
	} else {
		r.ParseForm()
		fmt.Println("username: ",r.Form["username"])
		fmt.Println("password: ",r.Form["password"])
	}
}


func upload	(w http.ResponseWriter,r *http.Request) {
	fmt.Println("method:",r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h,strconv.FormatInt(crutime,10))
		token := fmt.Sprintf("%x",h.Sum(nil))
		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w,token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file,handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println("err: ",err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w,"%v",handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename,os.O_WRONLY|os.O_CREATE,0666)
		if err != nil{
			fmt.Println("err: ",err)
			return
		}
		defer f.Close()
		io.Copy(f,file)

	}
}
func main() {
	http.HandleFunc("/",sayhelloName1)
	http.HandleFunc("/login",login)
	http.HandleFunc("/upload",upload)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ",err)
	}
}
