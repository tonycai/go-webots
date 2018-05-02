package main

import (
	"flag"
	"fmt"
)

/*
zizaike-free-wifi:workspace tonycai$ go run get_options.go --dir=abc
--filename set to "abc"
*/

type stringFlag struct {
	set   bool
	value string
}

func (sf *stringFlag) Set(x string) error {
	sf.value = x
	sf.set = true
	return nil
}

func (sf *stringFlag) String() string {
	return sf.value
}

var directory stringFlag

func init() {
	flag.Var(&directory, "dir", "The directory")
}

func main() {
	flag.Parse()
	if !directory.set {
		fmt.Println("--dir not set")
	} else {
		fmt.Printf("--dir set to %q\n", directory.value)
	}
}
