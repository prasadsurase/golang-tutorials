package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Main")
	wg.Add(1)

	go foo()

	wg.Add(1)

	go bar()

	wg.Wait()
}

func foo() {
	fmt.Println("foo")
	wg.Done()
}

func bar() {
	fmt.Println("bar")
	wg.Done()
}
