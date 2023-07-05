package main

import "fmt"

func main() {
	// May you declare the variable "var X = " or "var X type ="
	// Ex: var x = 1 or var a int = 1
	a := 10                    // int
	b := "golang"              // string
	c := 3.14                  // float64
	d := true                  // boolean
	e := "Hello"               // string
	f := `Do you like my hat?` // string
	g := 'M'                   // int32

	fmt.Printf("%v \t %T \n", a, a)  // 10       int
	fmt.Printf("%v \t %T  \n", b, b) // golang   string
	fmt.Printf("%v \t %T  \n", c, c) // 3.14     float64
	fmt.Printf("%v \t %T  \n", d, d) // true     bool
	fmt.Printf("%v \t %T  \n", e, e) // Hello    string
	fmt.Printf("%v \t %T  \n", f, f) // Do you like my hat?      string
	fmt.Printf("%v \t %T  \n", g, g) // 77       int32

	//Note https://stackoverflow.com/questions/41159169/why-not-use-v-to-print-int-and-string
	//Type %v verb means to use the default format which can be overridden.
}
