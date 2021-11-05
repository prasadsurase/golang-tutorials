package main

import "fmt"

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Prasad",
		friends: map[string]int{
			"A": 1,
			"B": 2,
			"C": 3,
		},
		favDrinks: []string{
			"beer",
			"Water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}
