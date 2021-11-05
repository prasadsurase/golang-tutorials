package main

import "fmt"

func main() {
	arr := []int{11, 22, 33, 44, 55, 66, 77, 88, 99, 100}

	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", arr)
}
