package main

import "fmt"

const metersToYards float64 = 1.01223

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters) // Enter meters
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}
