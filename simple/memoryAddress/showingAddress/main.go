package main

import "fmt"

func main() {
	a := 43

	fmt.Println("a- ", a)                    // 43
	fmt.Println("a's memory address - ", &a) // 0x... is memory address
	fmt.Printf("%v \n", &a)                  // 0x... is memory address
	fmt.Printf("%d \n", &a)                  // answer is int converted from memory address to integer.
}
