package main

import (
	"fmt"
)

var x = 42
var y = "James Bond"
var z = true

func main() {
	// String formatting
	s := fmt.Sprintf("Values are x: %v, y: %v, z: %v", x, y, z)
	fmt.Println(s)
}
