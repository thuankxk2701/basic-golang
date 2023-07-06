package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x...

	var b = &a

	fmt.Println(b)  //  Return value memory address
	fmt.Println(*b) //  Return value

	*b = 42        // b says, "The value at this address, change it to 42"
	fmt.Println(a) // 42

	// This is useful
	// We can pass a memory address instead of a bunch of values (we can pass a reference )
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// We don't have to pass around large amounts of data
	// We only have to pass around addresses

	// everything is PASS BY VALUE in go , btw etc...
	// When we pass a memory address, we are passing a value

}
