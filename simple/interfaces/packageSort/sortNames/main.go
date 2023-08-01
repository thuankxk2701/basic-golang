package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}

// https://golang.org/pkg/sort/#Sort
// https://golang.org/pkg/sort/#Interface

//package main
//
//import (
//"fmt"
//"sort"
//)
//
//func main() {
//	s := []int{5, 2, 6, 3, 1, 4} // unsorted
//	sort.Sort(sort.Reverse(sort.IntSlice(s)))
//	fmt.Println(s)
//}
