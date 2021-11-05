package main

import "fmt"

func main() {
	by := 1986
	for {
		if by <= 2021 {
			fmt.Println(by)
			by++
		} else {
			break
		}
	}
}
