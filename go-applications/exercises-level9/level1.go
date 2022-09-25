package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

const nr = 2

func main() {
	wg.Add(nr)

	for i := 0; i < nr; i++ {
		go func() {
			fmt.Println("Hello from a goroutine!")
			fmt.Println("# Current goroutines:", runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait()
}
