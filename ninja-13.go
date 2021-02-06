package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		for y := 0; y < 3; y++ {
			// ASCII Value
			fmt.Printf("%#U\n", i)
		}
		fmt.Printf("\n")

	}
}
