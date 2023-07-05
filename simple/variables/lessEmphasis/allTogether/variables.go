package main

import "fmt"

var a = "this is stored in the variable a"     // package scope
var b, c string = "stored in b", "stored in c" // package scope
var d string                                   // package scope

func main() {

	d = "stored in d" // declaration above; assignment here; package scope
	var e = 42        // function scope - subsequent variables have func scope:
	f := 43
	g := "stored in g"
	h, i := "stored in h", "stored in i"
	j, k, l, m := 44.7, true, false, 'm' // single quotes
	n := "n"                             // double quotes
	o := `o`                             // back ticks

	fmt.Println("a - ", a) // this is stored in the variable a
	fmt.Println("b - ", b) // stored in b
	fmt.Println("c - ", c) // stored in c
	fmt.Println("d - ", d) // stored in d
	fmt.Println("e - ", e) // 42
	fmt.Println("f - ", f) // 43
	fmt.Println("g - ", g) // stored in g
	fmt.Println("h - ", h) // stored in h
	fmt.Println("i - ", i) // stored in i
	fmt.Println("j - ", j) // 44.7
	fmt.Println("k - ", k) // true
	fmt.Println("l - ", l) // false
	fmt.Println("m - ", m) // 109
	fmt.Println("n - ", n) // n
	fmt.Println("o - ", o) // p
}
