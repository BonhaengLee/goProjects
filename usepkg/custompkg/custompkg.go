package custompkg

import (
	"fmt"
	"goProjects/usepkg/exinit"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

var Ratio int
var ttt int

const PI = 3.14
const pi2 = 3.1415

func PrintCustom() {
	fmt.Println("This is custom package!")
	exinit.PrintD()
}

func printcustom2() {
	fmt.Println("This is custom package2222222!")
}
