package main

import "fmt"

type foo int

var x foo

func main() {
	fmt.Printf("%T\t%v\n", x, x)
	x = 42
	fmt.Printf("%T\t%v\n", x, x)
}
