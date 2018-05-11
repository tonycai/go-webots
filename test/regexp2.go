package main

import "fmt"
import "regexp"

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println("FindAllString")
	elements := r.FindAllString("peach punch pinch", -1) // -1 unlimit
	fmt.Println(elements)
	for _, element := range elements {
		fmt.Println(element)
	}
	for i, element := range elements {
		fmt.Println(i, element)
	}

}
