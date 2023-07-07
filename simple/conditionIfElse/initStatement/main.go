package main

import "fmt"

func main() {
	b := true

	if food := "Chocolate"; b {
		fmt.Println(food) // Chocolate
	}
	// Note: init statement error invalid code
	//fmt.Println(food)

}
