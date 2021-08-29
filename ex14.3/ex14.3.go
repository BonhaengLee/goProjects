package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

// func ChangeData(arg Data) {
// 	arg.value = 999
// 	arg.data[100] = 999 // 100번째 인덱스에 999를 할당(왜?)
// }
// data를 변경시키고 싶다.
func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999 // 100번째 인덱스에 999를 할당(왜?)
}

func main() {
	var data Data

	// ChangeData(data) // 넣어준 Data타입의 인수(data)를 전달 받아서 내부의 프로퍼티를 변경해줄거라 예상(틀림)
	ChangeData(&data) // 넣어준 Data타입의 인수(data)를 전달 받아서 내부의 프로퍼티를 변경해줄거라 예상(틀림)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}
