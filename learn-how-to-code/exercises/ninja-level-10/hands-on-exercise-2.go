package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)

	// go bar(c) does not work
	bar(c)

	fmt.Println("bye")
}

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println("Value:", <-c)
}
