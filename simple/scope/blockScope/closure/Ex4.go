package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}

/*
Closure helps us limit the scope of variables used by multiple function
without closure, for two or more functions to have access to the same variable.
That variable would need to be package scope
*/
