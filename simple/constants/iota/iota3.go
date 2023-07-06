package main

import "fmt"

const (
	a3 = iota
	b3
	c3
)

const (
	d3 = iota
	e3
	f3
)

func main() {
	fmt.Println(a3) // 0
	fmt.Println(b3) // 1
	fmt.Println(c3) // 2
	fmt.Println(d3) // 0
	fmt.Println(e3) // 1
	fmt.Println(f3) // 2
}
