package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {

	b, err := ioutil.ReadFile("sample.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	//fmt.Println(b) // print the content as 'bytes'

	content := string(b) // convert content to a 'string'

	//reg := regexp.MustCompile(`(?i:^hello).*Go`)
	reg := regexp.MustCompile(`(?i:<img src=".+?")`) //忽略大小，非贪婪匹配
	//fmt.Printf("%q\n", reg.FindAllString(str, -1))
	elements := reg.FindAllString(content, 1)

	for i, element := range elements {
		fmt.Println(i, element)
		get_v(element)
	}

}

func get_v(content string) string {
	reg := regexp.MustCompile(`"(.+?)"`)
	elements := reg.FindAllStringSubmatch(content, -1)
	for i, element := range elements {
		fmt.Println(i, element[1])
	}
	return ""
}
