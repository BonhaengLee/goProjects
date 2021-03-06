package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request){ // resp를 write하는 w인자, 사용자의 요청 정보를 가진 r인자 
	fmt.Fprintf(w, "Hello World") // w인자가 있기 때문에 resp를 보낼 수 있다.
}

type fooHandler struct{}
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user) // r.Body는 인터페이스이므로 NewDecoder에 줄 수 있음
	if err!=nil { // json화의 에러 처리
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user) // 어떤 인터페이스르 받아서 json형태로 반환
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK) // 에러발생. 어떻게 body에 넣어서 보내냐? chrome rest client 이용
	fmt.Fprintf(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name")
	if name=="" {
		name="World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func NewHttpHandler() http.Handler{
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", barHandler)
	mux.Handle("/foo", &fooHandler{})

	return mux;
}