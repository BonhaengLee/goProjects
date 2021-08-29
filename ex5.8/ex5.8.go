package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin) // stdin는 변수명, :=는 선언대입문, os패키지안에 있는 스탠다드인풋을 나타냄
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b) // n 은 입력받은 개수,
	if err != nil {              // nil은 올바르지 않은 메모리 주소, 아무것도 아닌 것, nil이 아니라면 에러가 났다 라는것
		fmt.Println(err)       // err가 nil이 아니라서 이 err가 expected integer
		stdin.ReadString('\n') // 버퍼를 비워주고 다시 입력 기다림
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil { // nil은 올바르지 않은 메모리 주소, 아무것도 아닌 것, nil이 아니라면 에러가 났다 라는것
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
}
