package main

import "fmt"

func main() {

	cards := []string{"Get", "Post"}

	cards = append(cards, "Put")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {

	return "Five of diamonds"
}
