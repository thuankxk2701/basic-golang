package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	var answer []int
	for _, n := range numbers {
		if callback(n) {
			answer = append(answer, n)
		}
	}
	return answer
}

func main() {
	result := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(result) // [2 3 4]
}
