package main

import "fmt"

func main() {

	//create channel
	c := make(chan int)

	// create 10 go routines and each goroutine will send a value to the channel
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
		}()
	}

	//range over the channel and print the values
	for k := 0; k < 100; k++ {
		fmt.Println(<-c)
	}

}
