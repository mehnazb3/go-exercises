package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true
	// Formated Print statement
	fmt.Printf("Here are all values - x: %v,y: %v,z: %v  ", x, y, z)
	fmt.Printf("%T\n", x)
	// Single line print statement
	fmt.Println(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)
	fmt.Printf("%T\n", z)
	fmt.Println(z)
}
