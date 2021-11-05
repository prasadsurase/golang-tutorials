package main

import "fmt"

func main() {
	favSport := "surfing"
	switch favSport {
	case "surfing":
		fmt.Println("hi surfing")
	case "everything else":
		fmt.Println("bye everyone")
	}
}
