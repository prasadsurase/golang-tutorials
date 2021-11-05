package main

import "fmt"

const (
	a = iota + 2021
	b = iota + 2021
	c = iota + 2021
	d = iota + 2021
)

func main() {
	fmt.Println(a, b, c, d)
}
