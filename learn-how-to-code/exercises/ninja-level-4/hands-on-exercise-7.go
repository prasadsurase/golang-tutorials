package main

import "fmt"

func main() {
	// "James", "Bond", "Shaken, not stirred"
	// "Miss", "Moneypenny", "Helloooooo, James."

	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James"}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xs := [][]string{xs1, xs2}
	fmt.Println(xs)

	for i, v := range xs {
		fmt.Println("\nRecord:", i)
		for j, n := range v {
			fmt.Printf("Index: %v  Value: %v\n", j, n)
		}
	}
}
