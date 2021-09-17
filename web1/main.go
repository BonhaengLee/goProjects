package main

import (
	"goProjects/web1/myapp"
	"net/http"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler()) // 웹서버 구동하고 리퀘스트를 기다림
}