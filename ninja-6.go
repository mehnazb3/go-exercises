package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Println(x)
	fmt.Printf("%d\n", x)
	// Binary
	fmt.Printf("%b\n", x)
	// Hexa Decimal
	fmt.Printf("%#x\n", x)
}
