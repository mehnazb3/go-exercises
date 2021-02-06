package main

import "fmt"

func main() {
	year := 1993
	// For Loop without condition
	for {
		if year > 2020 {
			break
		}
		fmt.Println(year)
		year++
	}
}
