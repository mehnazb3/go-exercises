package main

import "fmt"

func main() {
	x := 42 == 42
	y := 10 <= 15
	z := 20 >= 15
	a := 20 != 15
	b := 20 > 15
	c := 15 < 20
	fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t\n", x, y, z, a, b, c)

}
