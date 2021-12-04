package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go Routines:", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup

	wg.Add(gs)
	for i := 1; i < gs; i++ {
		go func() {
			defer wg.Done()
			v := counter
			v++
			runtime.Gosched()
			counter = v
		}()
		fmt.Println("Go Routines:", runtime.NumGoroutine())
	}
	fmt.Println("Go Routines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
	wg.Wait()
}
