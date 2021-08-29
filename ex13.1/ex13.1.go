package main

import "fmt"

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func main() {
	var house House
	house.Address = "서울시 강남구 ,,,,"
	house.Size = 28
	house.Price = 10
	house.Category = "아파트"

	fmt.Printf("%v\n", house)
	fmt.Printf("주소:%s 사이즈:%d평 가격:%f억원 종류:%s\n", house.Address, house.Size, house.Price, house.Category)
	// float64타입이지만 10.000000(소수점 아래 여섯자리) 달고 싶지 않으면 %v라는 모든 타입의 서식문자 써야함
}
