package main

import (
	"fmt"
	"log"
	"os/exec"
)

func uuid() string {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s", out)
}

func main() {
	var out string = uuid()
	fmt.Printf("%s", out)
}

//generate file name
