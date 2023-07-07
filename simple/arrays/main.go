package main

import "fmt"

//func main() {
//
//	var x [58]string
//
//	for i := 65; i <= 122; i++ {
//		x[i-65] = string(i)
//	}
//	fmt.Println(x)
//	fmt.Println(x[42])
//}

//func main() {
//	var x [256]int
//
//	fmt.Println(len(x))
//	fmt.Println(x[42])
//	for i := 0; i < 256; i++ {
//		x[i] = i
//	}
//	for i, v := range x {
//		fmt.Printf("%v - %T - %b\n", v, v, v)
//		if i > 50 {
//			break
//		}
//	}
//}

//func main() {
//	var x [256]byte
//
//	fmt.Println(len(x))
//	fmt.Println(x[42])
//	for i := 0; i < 256; i++ {
//		x[i] = byte(i)
//	}
//	for i, v := range x {
//		fmt.Printf("%v - %T - %b\n", v, v, v)
//		if i > 50 {
//			break
//		}
//	}
//}

func main() {
	var x [256]string

	fmt.Println(len(x))
	fmt.Println(x[0])
	for i := 0; i < 256; i++ {
		x[i] = string(i)
	}
	for _, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, []byte(v))
	}
}
