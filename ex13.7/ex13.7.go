package main

import (
	"fmt"
	"unsafe"
)

// type User struct {
// 	A int8 // 1
// 	B int  // 8
// 	C int8 // 1
// 	D int  // 8
// 	E int8 // 1
// } // 19일 거라고 예상함

// 메모리 낭비를 막으려면 순서를 변경해라
type User struct {
	A int8 // 1
	C int8 // 1
	E int8 // 1
	B int  // 8
	D int  // 8
} // 3바이트 + 5바이트 패딩 + 16이니 24일 거라고 예상함

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user)) //안전하지 않은 함수들을 제공하는 패키지임
}
