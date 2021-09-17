package main

import "fmt"

// ! : 빈 인터페이스는 모든 타입을 다 받고 싶을 때 쓸 수 있다.
// 빈 인터페이스를 받아서 v.(type)으로 접근하면 들어오는 타입에 따라 switch할 수 있다.
func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not supported type %T:%v\n", t,t)
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("String")
	PrintVal(Student{15})
}