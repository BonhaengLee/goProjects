package main

import "fmt"

func main() {
	var x int8 = 4
	var y int8 = 64

	fmt.Printf("x:%08b x<<2:%08b x<<2: %d\n", x, x<<2, x<<2) // 2칸은 2^2를 의미한다. 따라서 4 * 4 = 16
	fmt.Printf("y:%08b y<<2:%08b y<<2: %d\n", y, y<<2, y<<2) // 64 * 2^2 = 256
}
