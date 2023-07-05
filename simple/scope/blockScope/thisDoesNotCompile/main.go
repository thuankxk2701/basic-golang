package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	foo()
}

func foo() {
	// no access to x
	// this does not compile
	// fmt.Println(x)
	// Note: Error because x not variable global not in the scope
}
