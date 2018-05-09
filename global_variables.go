package main

import (
	"fmt"
	"time"
)

var k int

func say(s string) {
	for i := 0; i < 15; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, k)
		k++
	}
}

func main() {
	go say("world")
	say("hello")
}
