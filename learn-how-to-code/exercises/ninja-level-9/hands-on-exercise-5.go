package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Go Routines:", runtime.NumGoroutine())

	var counter int64
	const gs = 15
	var wg sync.WaitGroup

	wg.Add(gs)
	for i := 1; i <= gs; i++ {
		go func() {
			defer wg.Done()

			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))

			// runtime.Gosched()
		}()
		// fmt.Println("Go Routines:\t", runtime.NumGoroutine())
	}
	// fmt.Println("Counter:\t", counter)
	wg.Wait()
}
