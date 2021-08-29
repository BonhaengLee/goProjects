package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for { // 초기문; 조건문이 없는 for문?, 아 while문이 따로 없고 for만 단독으로 적었을 때 무한루프구나
		fmt.Println("입력하세요")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil { // err에 뭔가 넘어왔다는 것은 에러가 발생했음을 의미한다. 원래는 nil이 할당되어 있다.
			fmt.Println("숫자를 입력하세요") // 쌍따옴표는 문자열, 홑따옴표는 문자 하나를 의미한다.

			// 키보드 버퍼를 지웁니다.
			stdin.ReadString('\n')
			continue
		}

		fmt.Printf("입력하신 숫자는 %d입니다.\n", number)
		if number%2 == 0 {
			break // 짝수 검사를 합니다.
		}
	}
}
