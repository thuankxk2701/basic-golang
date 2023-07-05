package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}

/*
Closure helps us limit the scope of variables used by multiple function
without closure, for two or more functions to have access to the same variable.
That variable would need to be package scope

anonymous function
a function without a name

func expression
assigning a func to a variable
*/
