package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
	TB = 1 << (iota * 10)
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB) // 10000000000
	fmt.Printf("%d\n", KB) // 1024
	fmt.Printf("%b\t", MB) // 100000000000000000000
	fmt.Printf("%d\n", MB) // 1048576
	fmt.Printf("%b\t", GB) // 1000000000000000000000000000000
	fmt.Printf("%d\n", GB) // 1073741824
	fmt.Printf("%b\t", TB) // 10000000000000000000000000000000000000000
	fmt.Printf("%d\n", TB) // 1099511627776
}
