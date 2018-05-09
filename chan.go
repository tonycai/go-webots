package main

import "fmt"

func main() {
	fmt.Println("Begin doing something!")
	c := make(chan int)
	var x int = 0
	go func() {
		c <- 10
		fmt.Printf("x=%d\n", x)
		fmt.Println("Doing something…")
		x = <-c
		c <- 1
		close(c)
	}()
	<-c
	fmt.Println("Done!")
}
