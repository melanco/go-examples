package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	const gr = 50
	var counter int64
	var wg sync.WaitGroup

	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			atomic.LoadInt64(&counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
