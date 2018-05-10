package main

import "fmt"

func main() {
	g := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case i := <-g:
				fmt.Println(i + 10)
			case <-quit:
				fmt.Println("B quit")
				return // exit
			}
		}
	}()
	for i := 0; i < 10; i++ {
		g <- i
	}
	quit <- true // 没办法等待B的退出只能Sleep
	fmt.Println("Main quit")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

/*
Output:
1
2
3
Main quit
*/
