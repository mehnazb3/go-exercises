package main

import "fmt"

type vehicle struct {
	doors []string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: []string{"door1", "door2"},
			color: "black",
		},
		fourWheel: true,
	}
	s1 := sedan{
		vehicle: vehicle{
			doors: []string{"door3", "door4"},
			color: "blue",
		},
		luxury: true,
	}
	fmt.Println(t1)
	fmt.Println(s1)
}
