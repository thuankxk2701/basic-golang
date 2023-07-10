package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array
	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	// greeting = append(greeting, "Suprabadham") //Append success because capacity is 5
	// greeting[3] = "suprabadham" Error : runtime error: index out of range [3] with length 3

	fmt.Println(greeting[2])
}
