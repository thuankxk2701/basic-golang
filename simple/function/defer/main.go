package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world()
	hello()
}

// Note: The defer keyword is used to delay the execution of a function until the surrounding function completes. The
// deferred function calls are executed in Last-In-First-Out order
