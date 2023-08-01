package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	//sort.StringSlice(s).Sort()
	//sort.Reverse(sort.StringSlice(s)) reverse string not working
	sort.Sort(sort.StringSlice(s))
	//sort.Sort(sort.Reverse(sort.StringSlice(s))) reverse string
	fmt.Println(s)

}

// You can use one in tow sort way!
// https://golang.org/pkg/sort/#Sort
