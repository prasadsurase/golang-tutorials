package main

import (
	"fmt"
)

func main() {
	f := counter()
	fmt.Println(f())
	fmt.Println(f())
	g := counter()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
}

func counter() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
