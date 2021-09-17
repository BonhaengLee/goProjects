package exinit

import (
	"fmt"
)

var (
	a = c + b
	b = f()
	c = f()
						d = 3
)

func init() {
	d++
	fmt.Println("exinit.init function", d)
}

func f() int {
	d++
	fmt.Println("exinit.f function", d)
	return d
}

func PrintD() {
	fmt.Println("d:", d)
}
