package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("\n Sum of all numbers: %v", sum(x...))
	fmt.Printf("\n Sum of all even numbers: %v", even(sum, x...))
	fmt.Printf("\n Sum of all odd numbers: %v", odd(sum, x...))

}

func sum(x ...int) int {
	total := 0

	for _, v := range x {
		total += v
	}

	return total
}

func even(f func(x ...int) int, y ...int) int {
	nos := []int{}

	for _, v := range y {
		if v%2 == 0 {
			nos = append(nos, v)
		}
	}
	return f(nos...)
}

func odd(f func(x ...int) int, y ...int) int {
	nos := []int{}

	for _, v := range y {
		if v%2 != 0 {
			nos = append(nos, v)
		}
	}
	return f(nos...)
}
