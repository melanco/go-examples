package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	const gr = 50
	counter := 0
	var wg sync.WaitGroup

	var mu sync.Mutex

	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
