package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
		close(c)
	}()
	//use comma ok idiom to check if there is still data on the channel.
	v, ok := <-c
	fmt.Println(v, ok)

	//Check again, should be false because channel has no more data on it.
	v, ok = <-c
	fmt.Println(v, ok)
}
