package main

import "fmt"

const (
	MasterRoom uint8 = 1 << iota
	LivingRoom
	BathRoom
	SmallRoom
)

func SetLight(rooms, room uint8) uint8 {
	return rooms | room
}

func ResetLight(rooms, room uint8) uint8 {
	return rooms &^ room
}

func IsLightOn(rooms, room uint8) bool {
	return rooms&room == room
}

func TurnLights(rooms uint8) {
	if IsLightOn(rooms, MasterRoom) {
		fmt.Println("안방에 불을 켠다")
	}
	if IsLightOn(rooms, LivingRoom) {
		fmt.Println("거실에 불을 켠다")
	}
	if IsLightOn(rooms, BathRoom) {
		fmt.Println("화장실에 불을 켠다")
	}
	if IsLightOn(rooms, SmallRoom) {
		fmt.Println("작은방에 불을 켠다")
	}

}

func main() {
	// 어떤 어떤 방에 불을 켜고 싶은지 미리 설정을 해놓고
	var rooms uint8 = 0
	// masterroom 자리의 비트 set하면 0000 0001
	rooms = SetLight(rooms, MasterRoom)
	// bathroom 자리의 비트 set하면 0000 0101
	rooms = SetLight(rooms, BathRoom)
	// smallroom 자리의 비트 set하면 0000 0111
	rooms = SetLight(rooms, SmallRoom)
	// smallroom 자리의 비트 reset하면 0000 0101
	rooms = ResetLight(rooms, SmallRoom)
	// 비트가 1인 것만 출력을 해라
	TurnLights(rooms)
}
