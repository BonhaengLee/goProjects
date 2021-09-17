package main

import (
	"net/http"
)



func main() {
	// // 동적 라우터, 왜 http가 아니라 이걸로 바꾸었는지?
	// mux := http.NewServeMux()
	// // 어떤 리퀘스트가 들어왔을 때 어떤 처리를 할 것인지에 대한 핸들러를 등록하는 HandlerFunc
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){ // resp를 write하는 w인자, 사용자의 요청 정보를 가진 r인자 
	// 	fmt.Fprintf(w, "Hello World") // w인자가 있기 때문에 resp를 보낼 수 있다.
	// })

	// mux.HandleFunc("/bar", barHandler)
	// // fooHandler 인스턴스를 만들어 넘겨줌
	// mux.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", mux) // 웹서버 구동하고 리퀘스트를 기다림
}