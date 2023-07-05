package main

import "fmt"

var x = 42 // Global

func main() {
	fmt.Println(x)
	foo()

}

func foo() {
	fmt.Println(x)

}
