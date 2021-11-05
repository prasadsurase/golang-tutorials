package main

import "fmt"

func main() {
	for i := 10; i <= 20; i++ {
		if i%4 != 0 {
			fmt.Println(i % 4)
		} else {
			fmt.Println("divisisble by 4")
		}
	}
}
