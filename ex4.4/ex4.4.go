package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5

	var c int = int(b)  // float >> int
	d := float64(a) * b // int >> float64

	var e int64 = 7
	f := a * int(e) // int64 >> int

	fmt.Println(a, b, c, d, e, f)
}
