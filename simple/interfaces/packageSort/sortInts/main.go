package main

import (
	"fmt"
	"sort"
)

func main() {

	n := []int{5, 3, 1, 6, 4, 2}
	sort.Ints(n)
	fmt.Println(n)
}
