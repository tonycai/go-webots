package main

import (
	"fmt"
	"regexp"
)

func main() {
	if regexp.MustCompile(`hello`).MatchString("Hello Regular Expression.") == true {
		fmt.Printf("Match ") // Will print 'Match' again
	} else {
		fmt.Printf("No match ")
	}
}
