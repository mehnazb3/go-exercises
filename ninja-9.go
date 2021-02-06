package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("%d\t%b\t%#x\t\n", x, x, x)
	// Left Byte shift
	y := x << 1
	fmt.Printf("%d\t%b\t%#x\t\n", y, y, y)
}
