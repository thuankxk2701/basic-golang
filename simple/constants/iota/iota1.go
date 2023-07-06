package main

import "fmt"

const (
	a1 = iota
	b1 = iota
	c1 = iota
)

func main() {
	fmt.Println(a1) // 0
	fmt.Println(b1) // 1
	fmt.Println(c1) // 2
}
