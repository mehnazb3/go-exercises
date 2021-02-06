package main

import "fmt"

func main() {
	// Modolus
	for i := 10; i <= 100; i++ {
		if i == 10 {
			fmt.Println("Staring")
		}
		fmt.Println(i % 4)
		if i == 100 {
			fmt.Println("Ending")
		}
	}
}
