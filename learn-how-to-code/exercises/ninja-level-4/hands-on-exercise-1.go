package main

import "fmt"

func main() {
	arr := [5]int{11, 22, 33, 44, 55}

	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", arr)
}
