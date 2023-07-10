package main

import "fmt"

func main() {
	var results []int
	fmt.Println(results) // result is []

	mySlice := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(mySlice)       // result is [a b c d e f]
	fmt.Println(mySlice[2:4])  // slicing a slice and result is [ c d ]
	fmt.Println(mySlice[2])    // index access; accessing by index and result is c
	fmt.Println("myString"[2]) // index access; accessing by index and result convert to string in code UTF-8

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}
	fmt.Print("[1:2] ")
	fmt.Println(greeting[1:2]) // result is [ Bonjour!]
	fmt.Print("[:2] ")
	fmt.Println(greeting[:2]) // result is [ Good morning! Bonjour! ]
	fmt.Print("[5:] ")
	fmt.Println(greeting[5:]) // result is 	[ Selamat pagi! Gutten morgen!]
	fmt.Print("[:] ")
	fmt.Println(greeting[:]) // result  is  [all]
}
