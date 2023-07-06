package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  //  Return value
	fmt.Println(&a) //  Return value memory address

	var b = &a

	fmt.Println(b)  //  Return value memory address
	fmt.Println(*b) //  Return value

	// The above code makes b a pointer to the memory address where an int is stores
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int
}
