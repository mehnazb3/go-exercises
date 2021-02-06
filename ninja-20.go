package main

import "fmt"

func main() {
	favSport := "Cricket"
	// Switch Statement
	switch favSport {
	case "Badminton":
		fmt.Println("Badminton")
	case "Cricket":
		fmt.Println("Cricket")
	default:
		fmt.Println("No favSport")
	}
}
