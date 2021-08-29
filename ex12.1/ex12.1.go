package main

import (
	"fmt"
)

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2} // go의 배열 선언은 특이하다. 중괄호로 묶어줌
	// t라는 변수에다가 같은 타입의 배열을 복사("=") 해준 거죠. :=가 아니라 =인 것은 var를 붙여주었기 때문임

	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}
}
