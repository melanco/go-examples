package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines BEFORE", runtime.NumGoroutine())

	var counter int64
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			//time.Sleep(time.Second)
			wg.Done()
		}()
		fmt.Println("Goroutines DURING", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("Counter", counter)
}
