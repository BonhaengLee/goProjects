package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32   // 4
	Score float64 // 8
}

func main() {
	user := User{23, 77.2}
	fmt.Println(unsafe.Sizeof(user)) //안전하지 않은 함수들을 제공하는 패키지임
}
