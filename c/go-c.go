package main

import (
	"os"
	"strings"
)

import "C"

//export FooBar
func FooBar() *C.char {
	args := os.Args
	return C.CString(strings.Join(args, " "))
}

func main() {}
