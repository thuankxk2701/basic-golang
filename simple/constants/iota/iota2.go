package main

import "fmt"

const (
	a2 = iota
	b2
	c2
)

func main() {
	fmt.Println(a2) // 0
	fmt.Println(b2) // 1
	fmt.Println(c2) // 2
}
