package main

import "fmt"

const (
	_ = iota
	//_ = iota  May you add multiple _= iota to increase value of variable
	b = iota * 10
	c = iota * 10
)

func main() {
	fmt.Println(b) // 10
	fmt.Println(c) // 20
}
