package main

import "fmt"

func max(x int) int {
	return 42 + x
}

func main() {
	max := max(7)
	fmt.Println(max) // Max is now the result, not the function result is 49
}

// Don't do this, because it will bad coding practice to shadow variables
