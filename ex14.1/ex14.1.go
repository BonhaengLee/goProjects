package main

import "fmt"

func main() {
	var a int = 500 // a는 일반 변수 값은 500
	var p *int      // p는 포인터 변수 미할당

	p = &a // p에 a의 메모리 주소 할당

	fmt.Printf("p의 값: %p\n", p)            // a의 메모리 주소
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p) // p가 가진 메모리주소에 접근하면 있는 값

	*p = 100                    // p가 가진 메모리주소로 접근해서 100을 할당
	fmt.Printf("a의 값: %d\n", a) // p가 가진 메모리주소(a)가 100이 할당되었으니 100이 나올듯
}
