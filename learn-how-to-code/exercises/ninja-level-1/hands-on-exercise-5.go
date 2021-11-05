package main

import "fmt"

type foo int

var x foo
var y int

func main() {
	fmt.Printf("%T\t%v\n", x, x)
	x = 42
	fmt.Printf("%T\t%v\n", x, x)

	y = int(x)
	fmt.Printf("%T\t%v\n", y, y)
}
