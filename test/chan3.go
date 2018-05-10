package main

import "fmt"

func main() {
	c := make(chan int, 1)
	go func() { c <- 1700 }()
	fmt.Println(<-c)
	go func() { c <- 2700 }()
	fmt.Println(<-c)
	for i := 0; i < 10; i++ {
		go func() { c <- 2700 + i }()
		fmt.Println(<-c)
	}
	go func() {
		c <- 3702
		fmt.Println("Done")
	}()
	c <- 5700
	j := <-c
	c <- j + 2000
	fmt.Println(<-c)
	close(c)
}
