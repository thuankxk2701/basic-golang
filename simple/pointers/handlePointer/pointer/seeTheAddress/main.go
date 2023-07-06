package main

import "fmt"

func zero(z *int) {
	fmt.Println(z) // 0x... like memory address of x
	*z = 0
}

func main() {
	x := 5
	fmt.Println(&x) // 0x...
	zero(&x)
	fmt.Println(x) // x is 0
}
