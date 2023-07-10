package main

import "fmt"

func main() {
	xs := []int{1, 3, 5, 69, 21}

	for i, value := range xs {
		fmt.Println(i, "-", value)
	}
}
