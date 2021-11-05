package main

import "fmt"

func main() {
	x := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}
	x()
	fmt.Println("hello world")
}
