package main

import "fmt"

type person struct {
	first_name   string
	last_name    string
	fav_icecream string
}

func main() {
	p1 := person{
		first_name:   "Test 1",
		last_name:    "Use 1",
		fav_icecream: "Black current",
	}
	p2 := person{
		first_name:   "Test 2",
		last_name:    "Use 12",
		fav_icecream: "Vanilla",
	}
	fmt.Println(p1)
	fmt.Println(p2)
}
