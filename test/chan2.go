package main

import (
	"fmt"
	"time"
)

func download(url string, ch chan int) {
	time.Sleep(5 * time.Second)
	fmt.Println(url)
	fmt.Println("done")
	ch <- 1 //return
}

func main() {
	urls := []string{
		"http://www.a.com/1.gzip",
		"http://www.a.com/2.gzip",
		"http://www.a.com/3.gzip",
		"http://www.a.com/4.gzip",
	}
	ch := make(chan int, 4)
	for _, v := range urls {
		go download(v, ch)
	}
	for i := 0; i < len(urls); i++ {
		<-ch
	}
	fmt.Println("over")
}
