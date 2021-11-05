package main

import "fmt"

func main() {
	x := []int{11, 22, 33, 44, 55}

	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
