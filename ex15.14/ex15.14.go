package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str1 := "Hello 월드"
	str2 := str1
	// stringHeader는 뭐지? reflect는 무슨 패키진지 모르겠고 unsafe는 안전하지 않은 메서드 Pointer인걸 보니 &strX의 주소를 가리키는 포인터 변수 생성?
	// 해설 : &str1을 raw pointer형태로 바꾼다음 그걸 다시 reflect.StringHeader로 바꾼단 얘기, struct인데 data, len으로 구성됨
	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
}
