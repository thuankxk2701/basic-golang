package main

import "fmt"

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	incrementA := wrapper()
	incrementB := wrapper()
	fmt.Println("A:", incrementA()) // 1
	fmt.Println("A:", incrementA()) // 2
	fmt.Println("B:", incrementB()) // 1
	fmt.Println("B:", incrementB()) // 2
	fmt.Println("B:", incrementB()) // 3
}

/*
Closure helps us limit the scope of variable used by multiple functions without closure
for two or more functions to have access to the same variable,
that variable would need to be package scope
*/
