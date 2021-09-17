package main

import "fmt"

type Attacker interface {
	Attack()
}

type Monster struct {
	Lv int
}

func (m *Monster) Attack() {
	fmt.Println("Monster Attack")
}

func DoAttack(att Attacker){
	if att != nil {
		att.Attack()
		// 여기서 몬스터를 꺼내오고 싶다면? 인터페이스 안의 인스턴스를 꺼내오고 싶다면? 바꾸고 싶은 타입으로 변환
		var monster *Monster
		monster = att.(*Monster)
		fmt.Println(monster.Lv)
		fmt.Println("%T",*monster)
	}
}

func main(){
	// var att Attacker
	// att.Attack()
	DoAttack(&Monster{20})
}