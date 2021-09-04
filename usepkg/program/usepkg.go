package main

import (
	"goProjects/usepkg/custompkg"
	"goProjects/usepkg/exinit"
	// "github.com/guptarohit/asciigraph"
	// "github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	exinit.PrintD()
	// s := custompkg.Student{} // 내부프로퍼티도 앞글자소문자면 private, 프로퍼티 개수도 맞춰야 바로 전달 가능
	// s.Name = "화랑"
	// s.Age = 32
	// fmt.Println(s.Name, s.Age)
	// fmt.Println(custompkg.PI)
	// custompkg.Ratio = 10
	// fmt.Println(custompkg.Ratio)
	// expkg.PrintSample()

	// expkg.PrintSample()

	// data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	// graph := asciigraph.Plot(data)
	// fmt.Println(graph)
}
