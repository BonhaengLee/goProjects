package main

import (
	"fmt"
	mrand "math/rand"
	_ "strings" // strings이 가진 요소의 부가 효과 때문에 임포트
)

func main() {
	fmt.Println(mrand.Int())
}
