package main

import (
	"fmt"
)

// Defining Custom Variable type
type customVariable int

var x customVariable
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	// Type conversion for underlying type
	y = int(x)
	fmt.Println(y)

}
