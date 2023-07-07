package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe"))
}

func greet(firstName, lastName string) string {
	return fmt.Sprint(firstName, lastName)
}
