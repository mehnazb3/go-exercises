package main

import (
	"fmt"
)

// Defining Custom Variable type
type customVariable int

var x customVariable

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
