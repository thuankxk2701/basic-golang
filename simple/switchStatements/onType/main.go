package main

import "fmt"

// Switch on types
// -- normally we switch on value of variable
// -- go allows you to switch on type of variable

type contact struct {
	greeting string
	name     string
}

// SwitchOnType  works with interfaces
// We'll learn more about interface later
func SwitchOnType(x interface{}) {
	switch x.(type) { // This is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	SwitchOnType(7)        // int
	SwitchOnType("McLeod") // string
	var t = contact{"Good to see you,", "Tim"}
	SwitchOnType(t)          // contact
	SwitchOnType(t.greeting) // string
	SwitchOnType(t.name)     // string
}
