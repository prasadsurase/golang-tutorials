package main

import "fmt"

type vehicle struct {
	doors int
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
	f150 := truck{
		vehicle: vehicle{
			doors: 4, color: "black",
		},
		fourWheel: true,
	}

	model3 := sedan{
		vehicle: vehicle{
			doors: 4, color: "silver",
		},
		luxury: true,
	}

	fmt.Println(f150)
	fmt.Println(model3)
}
