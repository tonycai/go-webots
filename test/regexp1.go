package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	/*
		r, err1 := regexp.Compile(`Jim|Tim`)

		if err1 == nil {
			fmt.Println(r.MatchString("Dickie, Tom and Tim")) // true
			fmt.Println(r.MatchString("Jimmy, John and Jim")) // true
		}
	*/

	r, err1 := regexp.Compile(`a$`) // without flag
	if err1 == nil {
		s := "atlanta\narkansas\nalabama\narachnophobia"
		//    01234567 890123456 78901234 5678901234567
		//                                            -
		res := r.FindAllStringIndex(s, -1)
		fmt.Printf("<%v>\n", res)
	}
	// 1 match
	// <[[37 38]]>
	t, err2 := regexp.Compile(`(?m)a$`) // with flag
	if err2 == nil {
		u := "atlanta\narkansas\nalabama\narachnophobia"
		//    01234567 890123456 78901234 5678901234567
		//          --                 --             -
		res2 := t.FindAllStringIndex(u, -1)
		fmt.Printf("<%v>", res2)
		// 3 matches
		// <[[6 7] [23 24] [37 38]]>
	}
	/*
	 */
}
