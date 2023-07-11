package main

import "fmt"

func main() {

	//myGreeting := map[string]string{}
	//myGreeting["Tim"] = "Good morning."
	//myGreeting["Jenny"] = "Bonjour."
	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	fmt.Println(myGreeting)
}
