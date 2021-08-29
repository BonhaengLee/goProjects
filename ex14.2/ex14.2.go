package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	var p1 *int = &a
	var p2 *int = &a // p1, p2는 a의 메모리 주소를 가짐
	var p3 *int = &b

	fmt.Printf("p1 == p2: %v\n", p1 == p2) // 둘의 메모리주소는 같을듯
	fmt.Printf("p2 == p3: %v\n", p2 == p3) // 둘은 당연히 다를듯
}
