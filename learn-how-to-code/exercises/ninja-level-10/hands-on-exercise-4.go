package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)

	receive(eve, odd, quit)
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("eve channel:", v)
		case v := <-o:
			fmt.Println("odd channel:", v)
		case v := <-q:
			fmt.Println("quit channel:", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	defer close(e)
	defer close(o)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}
