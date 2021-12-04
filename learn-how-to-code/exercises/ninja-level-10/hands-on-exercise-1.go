package main

import "fmt"

func main() {

	// using goroutine with closure.
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println("Meaning of life:", <-c)

	// using buffered channel
	c2 := make(chan int, 1)

	c2 <- 42

	fmt.Println("Meaning of life:", <-c2)

}
