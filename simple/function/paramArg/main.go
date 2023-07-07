package main

import "fmt"

func main() {

	greet("KeyXk")
	greetTowParams("Jane", "Doe")
}

func greet(name string) {
	fmt.Println(name)
}

func greetTowParams(firstName string, lastName string) { // You can declare params  ( firstName, lastName string )
	fmt.Println(firstName, lastName)
}
