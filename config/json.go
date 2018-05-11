package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	user := &User{Name: "Frank", Age: 30, Gender: "Male"}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	var str string
	str = string(b)
	fmt.Println(str)
}
