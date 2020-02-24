package main

import (
	"html/template"
	"net/http"
)

func welcome(res http.ResponseWriter,r *http.Request)  {
	files, _ := template.ParseFiles("view/index.html")
	files.Execute(res,nil)
}

func first(res http.ResponseWriter,r *http.Request){
	res.Write([]byte("first"))
}

func second(res http.ResponseWriter,r *http.Request){
	res.Write([]byte("second"))
}

func main(){
	server := http.Server{Addr: "localhost:8080"}
	//当发现url以/static开头把请求转发到指定路径
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))
	//配置控制器映射
	http.HandleFunc("/",welcome)
	http.HandleFunc("/first",first)
	http.HandleFunc("/second",second)
	server.ListenAndServe()
}
