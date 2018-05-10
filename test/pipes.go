package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if info.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gocowsay")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		input, _, err := reader.ReadLine() // ReadRune
		if err != nil && err == io.EOF {
			break
		}
		fmt.Printf("tony: %s \n", input)
	}
}
