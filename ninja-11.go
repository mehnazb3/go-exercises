package main

import "fmt"

const (
	// Iota to create sequence of year
	first  = 2014 + (1 * iota)
	second = 2014 + (1 * iota)
	third  = 2014 + (1 * iota)
	fourth = 2014 + (1 * iota)
)

func main() {
	fmt.Println(first, second, third, fourth)
}
