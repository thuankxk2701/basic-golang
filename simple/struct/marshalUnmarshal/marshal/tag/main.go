package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"Wisdom score"`
}

type personUnMarshal struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
}

func main() {
	p1 := person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))

	var pUn1 personUnMarshal
	fmt.Println(pUn1.First)
	fmt.Println(pUn1.Last)
	fmt.Println(pUn1.Age)

	bsUn := []byte(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.Unmarshal(bsUn, &pUn1)
	fmt.Println("--------------")
	fmt.Println(pUn1.First)
	fmt.Println(pUn1.Last)
	fmt.Println(pUn1.Age)
	fmt.Printf("%T \n", p1)
}
