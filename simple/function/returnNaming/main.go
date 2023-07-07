package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe"))
}

func greet(firstName, lastName string) (s string) {
	s = fmt.Sprint(firstName, lastName)
	return
}

/*
 IMPORTANT
 Avoid using named returns.

 Ocasionally named returns are useful. Read this article for more information:
 https://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html
*/
