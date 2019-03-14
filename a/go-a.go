package main

import (
	"os"
	"strings"
)

import "C"

//export Args
func Args() *C.char {
	args := os.Args
	return C.CString(strings.Join(args, " "))
}

func main() {}
